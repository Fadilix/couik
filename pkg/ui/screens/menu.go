package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/ui"
	"github.com/fadilix/couik/pkg/ui/modes"
)

type MenuScreen struct {
	Cursor  int
	Choices []string
}

func newMenuScreen() *MenuScreen {
	return &MenuScreen{
		Cursor: 0,
		Choices: []string{
			"15s",
			"30s",
			"60s",
			"120s",
			"quote",
			"words 10",
			"words 25",
		},
	}
}

func (s *MenuScreen) Update(m *ui.Model, msg tea.Msg) (ui.Screen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			if s.Cursor > 0 {
				s.Cursor--
			}
		case "right", "l":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter":
			selected := s.Choices[s.Cursor]

			var strategy modes.ModeStrategy

			switch selected {
			case "15s":
				strategy = modes.TimeMode{Duration: 15}
			case "30s":
				strategy = modes.TimeMode{Duration: 30}
			case "60s":
				strategy = modes.TimeMode{Duration: 60}
			case "quote":
				strategy = modes.QuoteMode{Language: database.English, Category: database.Mid}
			case "words 10":
				strategy = modes.WordMode{InitialWords: 10}
			case "words 25":
				strategy = modes.WordMode{InitialWords: 25}
			}
			return NewTypingScreen(strategy), nil
		}
	}
	return s, nil
}

func (s *MenuScreen) View(m *ui.Model) string {
	panic("unimplemented")
}
