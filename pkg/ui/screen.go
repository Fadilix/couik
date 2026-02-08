package ui

import tea "github.com/charmbracelet/bubbletea"

type Screen interface {
	Update(m *Model, msg tea.Msg) (Screen, tea.Cmd)
	View(m *Model) string
}
