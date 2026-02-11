package ui

import (
	"slices"
	"strconv"
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

	currentTime := defaultTMode.GetInitialTime()

	return Model{
		Session:         engine.NewSession(target),
		ProgressBar:     p,
		TimeLeft:        currentTime,
		initialTime:     currentTime,
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

// ApplyMode creates a new Model instance configured with the given strategy,
// preserving UI state from the current model
func (m Model) ApplyMode(mode modes.ModeStrategy) Model {
	config := mode.GetConfig()

	// create base model with target
	newModel := NewModel(config.Target)

	// apply mode strategy nd config
	newModel.Mode = mode
	newModel.CurrentLanguage = config.Language
	newModel.InitialWords = config.InitialWords
	newModel.QuoteType = config.Category

	// set time logic correctly
	newModel.TimeLeft = config.InitialTime
	newModel.initialTime = config.InitialTime

	// preserve ui state
	newModel.TerminalWidth = m.TerminalWidth
	newModel.TerminalHeight = m.TerminalHeight
	newModel.ProgressBar = m.ProgressBar
	newModel.Repo = m.Repo
	newModel.CustomDashboard = m.CustomDashboard

	return newModel
}

// GetModelFromMode does the same as the one above
// but just for the specific case of keytab
func (m *Model) GetModelFromMode(mod Model) Model {
	config := mod.Mode.GetConfig()
	model := NewModel(config.Target)

	model.Mode = mod.Mode
	model.CurrentLanguage = config.Language
	model.initialTime = mod.initialTime
	model.TimeLeft = mod.initialTime
	model.InitialWords = mod.InitialWords
	model.TerminalHeight = m.TerminalHeight
	model.TerminalWidth = m.TerminalWidth
	model.ProgressBar = m.ProgressBar
	return model
}

// CreateModeFromSelection parses the selection string and returns the appropriate ModeStrategy.
// It handles dynamic cases like "15s", "words 10"
func CreateModeFromSelection(selection string, lang database.Language) modes.ModeStrategy {
	// Time Mode
	if strings.HasSuffix(selection, "s") {
		seconds, _ := strconv.Atoi(strings.TrimSuffix(selection, "s"))
		return modes.NewTimeMode(modes.WithInitialTimeT(seconds), modes.WithLanguageT(lang))
	}

	// Word Mode
	if strings.HasPrefix(selection, "words ") {
		count, _ := strconv.Atoi(strings.TrimPrefix(selection, "words "))
		return modes.NewWordMode(modes.WithInitialWords(count), modes.WithLanguageW(lang))
	}

	// Quote Mode
	if selection == "quote" {
		return modes.NewQuoteMode(modes.WithLanguageQ(lang))
	}

	return nil
}
