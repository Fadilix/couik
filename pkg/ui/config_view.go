package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/cmd/couik/cli"
)

func (m Model) configView() string {
	config := cli.GetConfig()

	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	header := ViewHeaderStyle.Render(renderedLogo)

	keyStyle := lipgloss.NewStyle().Foreground(CatLavender).Bold(true).Width(15).Align(lipgloss.Left)
	valStyle := lipgloss.NewStyle().Foreground(CatText).Width(20).Align(lipgloss.Left)
	divider := lipgloss.NewStyle().Foreground(CatSurface).Render(" │ ")

	items := []struct{ k, v string }{
		{"Mode", config.Mode},
		{"DashASCII", config.DashboardASCII},
		{"Quote Type", config.QuoteType},
		{"Time", config.Time},
	}

	lines := []string{header, ""}
	for _, item := range items {
		row := lipgloss.JoinHorizontal(lipgloss.Center,
			keyStyle.Render(item.k),
			divider,
			valStyle.Render(item.v),
		)
		lines = append(lines, row)
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
