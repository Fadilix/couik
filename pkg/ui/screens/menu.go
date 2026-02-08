package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/pkg/ui"
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

func (s *MenuScreen) Update(m *ui.Model, msg tea.Cmd) (ui.Screen, tea.Cmd) {
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
			switch selected {
			// TODO: refactor the strategies and comeback
			// case "15s":
			// 	return m.GetDictionnaryModel(15), nil
			// case "30s":
			// 	return m.GetDictionnaryModel(30), nil
			// case "60s":
			// 	return m.GetDictionnaryModel(60), nil
			// case "quote":
			// 	return m.GetQuoteModel(), nil
			// case "words 10":
			// 	return m.GetDictionnaryModelWithWords(10), nil
			// case "words 25":
			// 	return m.GetDictionnaryModelWithWords(25), nil
			}
		}
	}
	return s, nil
}
