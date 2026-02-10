package ui

import (
	"fmt"

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

	lines := []string{}
	lines = append(lines, header, "\n")

	mode := fmt.Sprintf("%s %s\n", LabelStyle.Render("Mode"), ValueStyle.Render(config.Mode))
	dashASCII := fmt.Sprintf("%s %s\n", LabelStyle.Render("DashASCII"), ValueStyle.Render(config.DashboardASCII))
	quoteT := fmt.Sprintf("%s %s\n", LabelStyle.Render("Quote Type"), ValueStyle.Render(config.QuoteType))
	time := fmt.Sprintf("%s %s\n", LabelStyle.Render("Time"), ValueStyle.Render(config.Time))

	footer := HelpStyle.Render("[CTRL + R] return typing • [ESC] quit")

	lines = append(lines, mode, dashASCII, quoteT, time, footer)

	mui := lipgloss.JoinVertical(lipgloss.Center, lines...)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, mui)
}
