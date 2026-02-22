package ui

import (
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/cmd/couik/cli"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/internal/engine"
	"github.com/fadilix/couik/internal/storage"
	"github.com/fadilix/couik/pkg/network"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/components"
	"github.com/fadilix/couik/pkg/ui/core"
	"github.com/fadilix/couik/pkg/ui/modes"
)

type Model struct {
	Session        *engine.Session
	Repo           storage.HistoryRepository
	Quitting       bool
	ProgressBar    progress.Model
	TerminalWidth  int
	TerminalHeight int

	CurrentSelector components.Selector

	IsSelectingMode bool
	Mode            modes.ModeStrategy

	State core.SessionState

	TimeLeft    int
	initialTime int // Store the initial time duration for progress calculation
	Active      bool

	// quote mode selection
	QuoteType            database.QuoteCategory
	IsSelectingQuoteType bool

	// words
	InitialWords int

	// Cached config or use default
	CustomDashboard string
	CurrentLanguage database.Language

	// Cached chart for results view (prevents re-rendering shifts)
	CachedChart string

	// PR stats config
	PRs storage.Stats

	Client *network.Client
	// multiplayer
	Multiplayer      bool
	IsHost           bool
	PlayerName       string
	Mu               *sync.Mutex
	Countdown        int
	Players          map[string]*network.UpdatePayload
	LastDisconnected string
}

type NewModelOption func(*Model)

func WithMultiplayer() NewModelOption {
	return func(m *Model) {
		m.Multiplayer = true
		m.State = core.StateLobby
	}
}

func NewModel(target string, options ...NewModelOption) Model {
	p := progress.New(
		progress.WithSolidFill("#FFFFFF"),
		progress.WithWidth(20),
		progress.WithoutPercentage(),
	)
	p.Full = '━'
	p.Empty = '─'

	typingModes := []string{"15s", "30s", "60s", "120s", "quote", "w10", "w25"}

	defaultQT := database.Mid
	defaultInitTime := 30
	defaultDashboard := ""

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

	modeConfigContext := &core.ModeConfig{
		Target:       target,
		InitialWords: 50,
		Language:     cli.ParseConfigLang(config.Language),
		InitialTime:  defaultInitTime,
	}

	defaultTMode := modes.StringToMode(config.Mode, *modeConfigContext)

	currentTime := defaultTMode.GetInitialTime()

	loadedPrs := storage.LoadPRs()

	defaultState := core.StateTyping

	model := Model{
		Session:         engine.NewSession(target),
		ProgressBar:     p,
		TimeLeft:        currentTime,
		initialTime:     currentTime,
		Mode:            defaultTMode, // Default to quote mode
		QuoteType:       defaultQT,    // default to mid
		InitialWords:    50,
		CustomDashboard: defaultDashboard,
		Repo:            &storage.JSONRepository{},
		CurrentSelector: components.NewModeSelector(),
		PRs:             loadedPrs,
		Players:         make(map[string]*network.UpdatePayload),
		Mu:              &sync.Mutex{},
		State:           defaultState,
	}

	for _, option := range options {
		option(&model)
	}

	return model
}

func (m Model) Init() tea.Cmd {
	if m.Multiplayer && m.Client != nil {
		return WaitForNetworkMsg(m.Client)
	}
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
var newTarget strings.Builder

func (m Model) GetDictionnaryModelWithWords(words int, language database.Language) Model {
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

type ApplyModelOption func(*Model)

func WithSameQuote(target string) ApplyModelOption {
	return func(m *Model) {
		m.Session.Target = []rune(target)
	}
}

// ApplyMode creates a new Model instance configured with the given strategy,
// preserving UI state from the current model
func (m Model) ApplyMode(mode modes.ModeStrategy, options ...ApplyModelOption) Model {
	config := mode.GetConfig()

	newModel := NewModel(config.Target)

	newModel.Mode = mode
	newModel.CurrentLanguage = config.Language
	newModel.InitialWords = config.InitialWords
	newModel.QuoteType = config.Category

	newModel.TimeLeft = config.InitialTime
	newModel.initialTime = config.InitialTime

	newModel.TerminalWidth = m.TerminalWidth
	newModel.TerminalHeight = m.TerminalHeight
	newModel.ProgressBar = m.ProgressBar
	newModel.Repo = m.Repo
	newModel.CustomDashboard = m.CustomDashboard
	newModel.IsHost = m.IsHost
	newModel.Multiplayer = m.Multiplayer
	newModel.Client = m.Client
	newModel.PlayerName = m.PlayerName
	newModel.Players = m.Players
	newModel.Mu = m.Mu
	newModel.LastDisconnected = m.LastDisconnected

	for _, option := range options {
		option(&newModel)
	}

	return newModel
}

// GetModelFromMode does the same as the one above
// but just for the specific case of keytab and ctrl+r
func (m *Model) GetModelFromMode(mod Model) Model {
	var newMode modes.ModeStrategy
	switch mod.Mode.(type) {
	case *modes.QuoteMode:
		newMode = modes.NewQuoteMode(modes.WithLanguageQ(mod.CurrentLanguage), modes.WithCategoryQ(mod.QuoteType))
	case *modes.TimeMode:
		newMode = modes.NewTimeMode(modes.WithLanguageT(mod.CurrentLanguage), modes.WithInitialTimeT(mod.initialTime))
	case *modes.WordMode:
		newMode = modes.NewWordMode(modes.WithLanguageW(mod.CurrentLanguage), modes.WithInitialWords(mod.InitialWords))
	default:
		newMode = mod.Mode
	}
	return m.ApplyMode(newMode)
}

// CreateModeFromSelection parses the selection string and returns the appropriate ModeStrategy.
// It handles dynamic cases like "15s", "words 10"
func CreateModeFromSelection(selection string, lang database.Language) modes.ModeStrategy {
	if strings.HasSuffix(selection, "s") {
		seconds, _ := strconv.Atoi(strings.TrimSuffix(selection, "s"))
		return modes.NewTimeMode(modes.WithInitialTimeT(seconds), modes.WithLanguageT(lang))
	}

	if strings.HasPrefix(selection, "w") {
		count, _ := strconv.Atoi(strings.TrimPrefix(selection, "w"))
		return modes.NewWordMode(modes.WithInitialWords(count), modes.WithLanguageW(lang))
	}

	if selection == "quote" {
		return modes.NewQuoteMode(modes.WithLanguageQ(lang))
	}

	return nil
}

func WaitForNetworkMsg(c *network.Client) tea.Cmd {
	return func() tea.Msg {
		return <-c.NextMessage()
	}
}

func (m Model) PlayersConnected() int {
	return len(m.Players)
}
