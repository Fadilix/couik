package ui

import (
	"slices"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/cmd/couik/cli"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/internal/engine"
	"github.com/fadilix/couik/internal/storage"
	"github.com/fadilix/couik/pkg/typing"
)

type tickMsg time.Time

type sessionState int

// Typing mode
type TypingMode int

const (
	unselectedMode TypingMode = iota
	timedMode
	quoteMode
	wordMode
)

// Quote mode (small, mid, thicc)
type quoteType int

const (
	small quoteType = iota
	mid
	thicc
)

const (
	stateTyping sessionState = iota
	stateResults
)

type Model struct {
	Session        *engine.Session
	Repo           storage.HistoryRepository
	Quitting       bool
	ProgressBar    progress.Model
	TerminalWidth  int
	TerminalHeight int

	// for mode selecting
	IsSelectingMode bool
	Cursor          int
	Choices         []string
	Mode            TypingMode

	// state
	State sessionState

	// timer
	timeLeft    int
	initialTime int // Store the initial time duration for progress calculation
	Active      bool

	// quote mode selection
	QuoteType            quoteType
	QuoteTypeCursor      int
	QuoteTypeChoices     []string
	IsSelectingQuoteType bool

	// words
	InitialWords int

	// User defaults
	// config cli.Config
	CustomDashboard string
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
	qType := []string{"small", "mid", "thicc"}

	defaultTMode := quoteMode
	defaultQT := mid
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

	switch config.Mode {
	case "quote":
		defaultTMode = quoteMode
	case "words":
		defaultTMode = wordMode
	case "time":
		defaultTMode = timedMode
	}

	return Model{
		Session:          engine.NewSession(target),
		ProgressBar:      p,
		Choices:          typingModes,
		timeLeft:         defaultInitTime,
		initialTime:      defaultInitTime,
		Mode:             defaultTMode, // Default to quote mode since we start with a random quote
		QuoteType:        defaultQT,    // default to mid
		InitialWords:     50,
		QuoteTypeChoices: qType,
		CustomDashboard:  defaultDashboard,
		Repo:             &storage.JSONRepository{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// func (m Model) CalculateTypingSpeed() float64 {
// 	// Guard: return 0 if not in results state to avoid division by zero
// 	if m.State != stateResults {
// 		return 0
// 	}
//
// 	var duration time.Duration
// 	var correctChars int
//
// 	if m.Mode == timedMode {
// 		duration = time.Duration(m.initialTime) * time.Second
// 		correctChars = game.CountCorrect(m.Results[:m.Index])
// 	} else {
// 		duration = m.EndTime.Sub(m.StartTime)
// 		correctChars = game.CountCorrect(m.Results)
// 	}
//
// 	wpm := ((float64(correctChars)) * (60 / duration.Seconds())) / 5
// 	return wpm
// }

// func (m Model) CalculateRawTypingSpeed() float64 {
// 	// Guard: return 0 if not in results state to avoid division by zero
// 	if m.State != stateResults {
// 		return 0
// 	}
//
// 	var duration time.Duration
// 	var correctChars, incorrectChars int
//
// 	if m.Mode == timedMode {
// 		duration = time.Duration(m.initialTime) * time.Second
// 		correctChars = game.CountCorrect(m.Results[:m.Index])
// 		incorrectChars = game.CountIncorrect(m.Results[:m.Index])
// 	} else {
// 		duration = m.EndTime.Sub(m.StartTime)
// 		correctChars = game.CountCorrect(m.Results)
// 		incorrectChars = game.CountIncorrect(m.Results)
// 	}
//
// 	wpm := ((float64(correctChars) + float64(incorrectChars)) * (60 / duration.Seconds())) / 5
// 	return wpm
// }
//
// func (m Model) CalculateAccuracy() float64 {
// 	var correctChars int
// 	var totalChars int
//
// 	if m.Mode == timedMode {
// 		correctChars = game.CountCorrect(m.Results[:m.Index])
// 		totalChars = m.Index // Use characters typed for timed mode
// 	} else {
// 		correctChars = game.CountCorrect(m.Results)
// 		totalChars = len(m.Target)
// 	}
//
// 	if totalChars == 0 {
// 		return 0
// 	}
//
// 	acc := (float64(correctChars-m.BackSpaceCount) / float64(totalChars)) * 100
// 	return acc
// }

func (m Model) GetQuoteModel() Model {
	quote := typing.GetQuoteUseCase(database.English, database.Mid)
	target := quote.Text

	newModel := NewModel(target)
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.timeLeft = 15
	newModel.Mode = quoteMode

	return newModel
}

func (m Model) GetDictionnaryModel(duration int) Model {
	newTarget := typing.GetDictionnary()

	newModel := NewModel(newTarget)
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.timeLeft = duration
	newModel.initialTime = duration
	newModel.Mode = timedMode

	return newModel
}

// GetDictionnaryModelWithWords creates a model with custom words length
// for word mode typing tests
func (m Model) GetDictionnaryModelWithWords(words int) Model {
	var newTarget strings.Builder
	dictionnary := typing.GetDictionnary()

	wordCount := 0

	for i := range dictionnary {
		if string(dictionnary[i]) == " " {
			wordCount++
			if wordCount == words {
				break
			}
		}
		newTarget.WriteString(string(dictionnary[i]))
	}
	newModel := NewModel(newTarget.String())
	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.Mode = wordMode
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
	newModel.Mode = wordMode
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
	newModel.timeLeft = initialTime
	newModel.Mode = timedMode
	newModel.InitialWords = len([]rune(target))
	newModel.initialTime = initialTime

	return newModel
}

func (m Model) GetModelWithQuoteType(option string) Model {
	var category database.Category

	switch option {
	case "mid":
		category = database.Mid
	case "thicc":
		category = database.Thicc
	default:
		category = database.Small
	}

	target := typing.GetQuoteUseCase(database.English, category)
	quote := target.Text

	newModel := NewModel(quote)

	newModel.TerminalHeight = m.TerminalHeight
	newModel.TerminalWidth = m.TerminalWidth
	newModel.ProgressBar.Width = m.ProgressBar.Width
	newModel.Mode = quoteMode

	return newModel
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
