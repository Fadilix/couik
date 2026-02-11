package ui

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/cmd/couik/cli"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/internal/engine"
	"github.com/fadilix/couik/internal/storage"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/components"
	"github.com/fadilix/couik/pkg/ui/core"
	"github.com/fadilix/couik/pkg/ui/modes"
)

// Typing mode
// type TypingMode int

// const (
// 	unselectedMode TypingMode = iota
// 	timedMode
// 	quoteMode
// 	wordMode
// )

// Quote mode (small, mid, thicc)

type Model struct {
	Session        *engine.Session
	Repo           storage.HistoryRepository
	Quitting       bool
	ProgressBar    progress.Model
	TerminalWidth  int
	TerminalHeight int

	// mode selecting
	CurrentSelector components.Selector

	// for mode selecting
	IsSelectingMode bool
	// Cursor          int
	// Choices         []string
	Mode modes.ModeStrategy

	// state
	State core.SessionState

	// timer
	TimeLeft    int
	initialTime int // Store the initial time duration for progress calculation
	Active      bool

	// quote mode selection
	QuoteType database.QuoteCategory
	// QuoteTypeCursor      int
	// QuoteTypeChoices     []string
	IsSelectingQuoteType bool

	// words
	InitialWords int

	// User defaults
	// config cli.Config
	CustomDashboard string
	CurrentLanguage database.Language
}

func NewModel(target string) Model {
	p := progress.New(
		progress.WithSolidFill("#FFFFFF"),
		progress.WithWidth(20),
		progress.WithoutPercentage(),
	)
	p.Full = '━'
	p.Empty = '─'

	typingModes := []string{"15s", "30s", "60s", "120s", "quote", "words 10", "words 25"}
	// qType := []string{"small", "mid", "thicc"}

	// default typing mode

	defaultQT := database.Mid
	defaultInitTime := 30
	defaultDashboard := ""

	// load the user config
	config := cli.GetConfig()

	if database.FileExists(config.DashboardASCII) {
		dashboard, _ := cli.GetTextFromFile(config.DashboardASCII)
		defaultDashboard = dashboard
	}

	if !slices.Contains(typingModes, config.Time) {
		switch config.Time {
		case "s":
			defaultInitTime = 15
		case "30s":
			defaultInitTime = 30
		case "60s":
			defaultInitTime = 60
		case "120s":
			defaultInitTime = 120
		}
	}

	// defaultTMode := modes.NewQuoteMode()

	modeConfigContext := &core.ModeConfig{
		Target:       target,
		InitialWords: 50,
		Language:     cli.ParseConfigLang(config.Language),
		InitialTime:  defaultInitTime,
	}

	defaultTMode := modes.StringToMode(config.Mode, *modeConfigContext)

	return Model{
		Session:         engine.NewSession(target),
		ProgressBar:     p,
		TimeLeft:        defaultInitTime,
		initialTime:     defaultTMode.GetInitialTime(),
		Mode:            defaultTMode, // Default to quote mode since we start with a random quote
		QuoteType:       defaultQT,    // default to mid
		InitialWords:    50,
		CustomDashboard: defaultDashboard,
		Repo:            &storage.JSONRepository{},
		CurrentSelector: components.NewModeSelector(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) GetQuoteModel() Model {
	quote := typing.GetQuoteUseCase(m.CurrentLanguage, database.Mid)
	target := quote.Text

	newModel := NewModel(target)
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.TimeLeft = 15
	newModel.Mode = modes.NewQuoteMode()

	return newModel
}

func (m Model) GetDictionnaryModel(duration int) Model {
	newTarget := typing.GetDictionnary(m.CurrentLanguage)

	newModel := NewModel(newTarget)
	newModel.CurrentLanguage = m.CurrentLanguage
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.TimeLeft = duration
	newModel.initialTime = duration
	newModel.Mode = modes.NewTimeMode(modes.WithLanguageT(m.CurrentLanguage), modes.WithTargetT(newTarget))

	return newModel
}

// GetDictionnaryModelWithWords creates a model with custom words length
// for word mode typing tests
func (m Model) GetDictionnaryModelWithWords(words int, language database.Language) Model {
	var newTarget strings.Builder
	dictionnary := typing.GetDictionnary(language)

	wordCount := 0

	for _, r := range dictionnary {
		if r == ' ' {
			wordCount++
			if wordCount == words {
				break
			}
		}
		newTarget.WriteRune(r)
	}
	newModel := NewModel(newTarget.String())
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.Mode = modes.NewWordMode(modes.WithInitialWords(words), modes.WithLanguageW(language))
	newModel.InitialWords = words

	return newModel
}

// GetModelWithCustomTarget creates a model with custom target
// for word mode typing tests
func (m Model) GetModelWithCustomTarget(target string) Model {
	newModel := NewModel(target)
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.Mode = modes.NewWordMode(modes.WithTargetW(target), modes.WithInitialWords(len([]rune(target))))
	newModel.InitialWords = len([]rune(target))

	return newModel
}

// GetTimeModelWithCustomTarget creates a model with custom target
// for time mode typing tests
func (m Model) GetTimeModelWithCustomTarget(initialTime int, target string) Model {
	newModel := NewModel(target)
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.TimeLeft = initialTime
	newModel.Mode = modes.NewTimeMode(modes.WithTargetT(target))
	newModel.InitialWords = len([]rune(target))
	newModel.initialTime = initialTime

	return newModel
}

func (m Model) GetModelWithQuoteType(quoteOption string) Model {
	var category database.QuoteCategory

	switch quoteOption {
	case "mid":
		category = database.Mid
	case "thicc":
		category = database.Thicc
	default:
		category = database.Small
	}

	target := typing.GetQuoteUseCase(m.CurrentLanguage, category)
	quote := target.Text

	newModel := NewModel(quote)
	newModel.CurrentLanguage = m.CurrentLanguage
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.Mode = modes.NewQuoteMode(modes.WithCategoryQ(category), modes.WithTargetQ(target.Text), modes.WithLanguageQ(m.CurrentLanguage))

	return newModel
}

func (m *Model) GetTimeLeft() int {
	return m.TimeLeft
}

func (m *Model) SetTimeLeft(t int) {
	m.TimeLeft = t
}

func (m *Model) Deactivate() {
	m.Active = false
}

func (m *Model) SetState(s core.SessionState) {
	m.State = s
}

func (m *Model) GetSession() *engine.Session {
	return m.Session
}

func (m *Model) GetModelFromMode(mode modes.ModeStrategy) Model {
	modeConfig := mode.GetConfig()

	model := NewModel(modeConfig.Target)
	model.InitialWords = m.InitialWords
	model.CurrentLanguage = modeConfig.Language
	model.Mode = mode
	model.TerminalHeight = m.TerminalHeight
	model.TerminalWidth = m.TerminalWidth
	model.ProgressBar = m.ProgressBar

	return model
}
