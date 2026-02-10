package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type cmdItem struct {
	key  string
	desc string
}

func (m Model) commandPaletteView() string {
	cmdPalette := []cmdItem{
		{"ESC", "Quit"},
		{"CTRL + R", "Refresh test / return typing)"},
		{"CTRL + L", "Restart the same test"},
		{"CTRL + E", "Choose quote type"},
		{"SHIFT + TAB", "Choose typing mode"},
		{"TAB", "Restart (when on results page)"},
		{"CTRL + P", "Show this palette"},
	}

	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	header := ViewHeaderStyle.Render(renderedLogo)

	lines := []string{}
	lines = append(lines, header, "\n")

	for _, value := range cmdPalette {
		cmd := LabelStyle.Render(value.key)
		desc := ValueStyle.Render(value.desc)

		line := fmt.Sprintf("%s %s\n", cmd, desc)
		lines = append(lines, line)
	}

	mui := lipgloss.JoinVertical(lipgloss.Center, lines...)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, mui)
}
