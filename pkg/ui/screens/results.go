package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/internal/engine"
	"github.com/fadilix/couik/pkg/ui"
)

type ResultsScreen struct {
	Session  *engine.Session
	ModeTime int
}

func NewResultsScreen(session *engine.Session, modeTime int) *ResultsScreen {
	return &ResultsScreen{
		Session:  session,
		ModeTime: modeTime,
	}
}

func (s *ResultsScreen) Update(m *ui.Model, msg tea.Msg) (ui.Screen, tea.Cmd) {
	// Save on enter (or immediately in Init, but here is fine for now)
	// We'll just listen for navigation keys
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			m.Quitting = true
			return s, tea.Quit
		case tea.KeyTab:
			return NewMenuScreen(), nil
		}
	}
	return s, nil
}

func (s *ResultsScreen) View(m *ui.Model) string {
	wpm := s.Session.CalculateTypingSpeed()
	acc := s.Session.CalculateAccuracy()

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center,
			ui.HeaderStyle.Render("Results"),
			ui.StatsStyle.Render(fmt.Sprintf("WPM: %.2f", wpm)),
			ui.StatsStyle.Render(fmt.Sprintf("Accuracy: %.2f%%", acc)),
			"\n",
			ui.PendingStyle.Render("[TAB] Menu • [ESC] Quit"),
		),
	)
}
