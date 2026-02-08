package screens

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/internal/engine"
	"github.com/fadilix/couik/pkg/ui"
	"github.com/fadilix/couik/pkg/ui/modes"
)

type TypingScreen struct {
	Session     *engine.Session
	InitialTime int
	TimeLeft    int
}

func NewTypingScreen(strategy modes.ModeStrategy) *TypingScreen {
	target := strategy.GetTarget()
	initialTime := strategy.GetInitialTime()

	return &TypingScreen{
		Session:     engine.NewSession(target),
		InitialTime: initialTime,
		TimeLeft:    initialTime,
	}
}

func (s *TypingScreen) Update(m *ui.Model, msg tea.Msg) (ui.Screen, tea.Cmd) {
	switch msg := msg.(type) {
	case ui.TickMsg:
		if s.InitialTime > 0 && s.Session.Started && !s.Session.IsFinished() {
			s.TimeLeft--
			if s.TimeLeft <= 0 {
				s.Session.EndTime = time.Now()
				return NewResultsScreen(s.Session, s.InitialTime), nil
			}
			return s, ui.Tick()
		}

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			m.Quitting = true
			return s, tea.Quit
		case tea.KeyBackspace:
			s.Session.BackSpace()
		case tea.KeyRunes, tea.KeySpace:
			char := msg.String()
			if msg.Type == tea.KeySpace {
				char = " "
			}

			s.Session.Type(char)

			// Start timer on first input
			if s.InitialTime > 0 && s.Session.Index == 1 && s.TimeLeft == s.InitialTime {
				return s, ui.Tick()
			}

			if s.Session.IsFinished() {
				// s.Session.IsFinished() now sets EndTime in your implementation, so we are good
				return NewResultsScreen(s.Session, s.InitialTime), nil
			}
		}
	}
	return s, nil
}

func (s *TypingScreen) View(m *ui.Model) string {
	// Simple text rendering for brevity
	var sb strings.Builder
	for i, r := range s.Session.Target {
		char := string(r)
		if i < s.Session.Index {
			if s.Session.Results[i] {
				sb.WriteString(ui.CorrectStyle.Render(char))
			} else {
				sb.WriteString(ui.WrongStyle.Render(char))
			}
		} else if i == s.Session.Index {
			sb.WriteString(ui.HighlightStyle.Render(char))
		} else {
			sb.WriteString(ui.PendingStyle.Render(char))
		}
	}

	// Add word wrapping or use the existing complex logic from old view.go if needed
	// For now, let's keep it simple to verify it works
	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center,
			ui.StatsStyle.Render("Typing..."),
			"\n",
			sb.String(),
		),
	)
}
