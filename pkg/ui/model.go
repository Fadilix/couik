package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Target         string
	Results        []bool
	Index          int
	StartTime      time.Time
	Started        bool
	Quitting       bool
	ProgressBar    progress.Model
	TerminalWidth  int
	TerminalHeight int
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
