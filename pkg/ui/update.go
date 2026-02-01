package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/pkg/typing"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height

		m.ProgressBar.Width = msg.Width - 20
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Quitting = true
			return m, tea.Quit

		case tea.KeyBackspace:
			if m.Index > 0 {
				m.Index--
			}

		case tea.KeyTab:
			if m.State == stateResults {
				newTarget := typing.GetRandomQuote()
				return NewModel(newTarget), nil
			}

		case tea.KeyRunes, tea.KeySpace:
			if m.State == stateResults {
				return m, nil
			}
			if !m.Started {
				m.StartTime = time.Now()
				m.Started = true
			}

			if m.Index < len(m.Target) {
				typedChar := msg.String()
				if msg.Type == tea.KeySpace {
					typedChar = " "
				}
				m.Results[m.Index] = typedChar == string(m.Target[m.Index])
				m.Index++
			}

			if m.Index >= len(m.Target) {
				m.State = stateResults
				m.EndTime = time.Now()
				// return m, tea.Quit
			}

		}
	}
	return m, nil
}
