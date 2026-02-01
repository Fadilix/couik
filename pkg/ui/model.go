package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/internal/game"
)

type sessionState int

const (
	stateTyping sessionState = iota
	stateResults
)

type Model struct {
	Target         string
	Results        []bool
	Index          int
	Started        bool
	Quitting       bool
	ProgressBar    progress.Model
	TerminalWidth  int
	TerminalHeight int

	// for better accuracy calculation
	BackSpaceCount int
	IsError        bool

	// state
	State sessionState

	// timer
	StartTime time.Time
	EndTime   time.Time
}

func NewModel(target string) Model {
	p := progress.New(
		progress.WithSolidFill("#FFFFFF"),
		progress.WithWidth(20),
		progress.WithoutPercentage(),
	)
	p.Full = '━'
	p.Empty = '─'

	return Model{
		Target:      target,
		Results:     make([]bool, len(target)),
		ProgressBar: p,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) CalculateTypingSpeed() float64 {
	var duration time.Duration

	if m.State == stateResults {
		duration = m.EndTime.Sub(m.StartTime)
	}

	correctChars := game.CountCorrect(m.Results)
	wpm := ((float64(correctChars)) * (60 / duration.Seconds())) / 5
	return wpm
}

func (m Model) CalculateRawTypingSpeed() float64 {
	var duration time.Duration

	if m.State == stateResults {
		duration = m.EndTime.Sub(m.StartTime)
	}

	correctChars := game.CountCorrect(m.Results)
	incorrectChars := game.CountIncorrect(m.Results)

	wpm := ((float64(correctChars) + float64(incorrectChars)) * (60 / duration.Seconds())) / 5
	return wpm
}

func (m Model) CalculateAccuracy() float64 {
	correctChars := game.CountCorrect(m.Results)
	acc := (float64(correctChars-m.BackSpaceCount) / float64(len(m.Target))) * 100
	return acc
}
