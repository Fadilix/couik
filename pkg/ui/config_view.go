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

	header := lipgloss.NewStyle().Foreground(CatMauve).Bold(true).Render(renderedLogo)

	LabelStyle := lipgloss.NewStyle().Foreground(CatMauve).Width(15).Align(lipgloss.Left)

	ValueStyle := lipgloss.NewStyle().Foreground(CatSubtext).Bold(true).Width(30).Align(lipgloss.Right)

	lines := []string{}
	lines = append(lines, header, "\n")

	mode := fmt.Sprintf("%s %s\n", LabelStyle.Render("Mode"), ValueStyle.Render(config.Mode))
	dashASCII := fmt.Sprintf("%s %s\n", LabelStyle.Render("DashASCII"), ValueStyle.Render(config.DashboardASCII))
	quoteT := fmt.Sprintf("%s %s\n", LabelStyle.Render("Quote Type"), ValueStyle.Render(config.QuoteType))
	time := fmt.Sprintf("%s %s\n", LabelStyle.Render("Time"), ValueStyle.Render(config.Time))

	helpStyle := lipgloss.NewStyle().Foreground(CatOverlay).MarginTop(1)
	footer := helpStyle.Render("[CTRL + R] return typing • [ESC] quit")

	lines = append(lines, mode, dashASCII, quoteT, time, footer)

	mui := lipgloss.JoinVertical(lipgloss.Center, lines...)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, mui)
}
