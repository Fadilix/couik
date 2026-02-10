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

	header := lipgloss.NewStyle().Foreground(CatMauve).Bold(true).Render(renderedLogo)

	LabelStyle := lipgloss.NewStyle().Foreground(CatMauve).Width(15).Align(lipgloss.Left)

	ValueStyle := lipgloss.NewStyle().Foreground(CatSubtext).Bold(true).Width(30).Align(lipgloss.Right)

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
