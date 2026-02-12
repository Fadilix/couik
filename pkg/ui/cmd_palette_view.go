package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type cmdItem struct {
	key  string
	desc string
}

func (m Model) commandPaletteView() string {
	cmdPalette := []cmdItem{
		{"ESC", "Quit"},
		{"CTRL + R", "Refresh test/return to typing view"},
		{"CTRL + L", "Restart the same test"},
		{"CTRL + E", "Choose quote type"},
		{"SHIFT + TAB", "Choose typing mode"},
		{"TAB", "Restart (when on results page)"},
		{"CTRL + P", "Show this palette"},
		{"CTRL + N", "Swtich to French/English"},
		{"CTRL + G", "Show user config"},
	}

	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	header := ViewHeaderStyle.Render(renderedLogo)

	keyStyle := lipgloss.NewStyle().Foreground(CatLavender).Bold(true).Width(15).Align(lipgloss.Left)
	descStyle := lipgloss.NewStyle().Foreground(CatSubtext).Width(35).Align(lipgloss.Left)
	divider := lipgloss.NewStyle().Foreground(CatSurface).Render(" │ ")

	lines := []string{}
	lines = append(lines, header, "")

	for _, value := range cmdPalette {
		line := lipgloss.JoinHorizontal(lipgloss.Center,
			keyStyle.Render(value.key),
			divider,
			descStyle.Render(value.desc),
		)
		lines = append(lines, line)
	}

	footerKey := lipgloss.NewStyle().Foreground(CatLavender).Bold(true)
	footerDesc := lipgloss.NewStyle().Foreground(CatOverlay)
	footer := lipgloss.JoinHorizontal(lipgloss.Center,
		footerKey.Render("ctrl+r "), footerDesc.Render("back"),
		footerDesc.Render("  •  "),
		footerKey.Render("esc "), footerDesc.Render("quit"),
	)
	lines = append(lines, "", footer)

	mui := lipgloss.JoinVertical(lipgloss.Center, lines...)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, mui)
}
