package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/pkg/ui/modes"
)

func (m Model) resultsView() string {
	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	// Logo Section
	header := ViewHeaderStyle.Render(renderedLogo)

	var formattedTime string
	_, isTimeMode := m.Mode.(*modes.TimeMode)

	if !isTimeMode {
		optionalTime := FormatTime(int(m.Session.EndTime.Sub(m.Session.StartTime).Seconds()))
		formattedTime = fmt.Sprintf("%s %s", LabelStyle.Render("Time:"), ValueStyle.Render(optionalTime))
	}

	// Build the stats block
	statsBox := lipgloss.JoinVertical(lipgloss.Left,
		StatsTitleStyle.Render("SESSION PERFORMANCE"),
		fmt.Sprintf("%s %s", LabelStyle.Render("Speed:"), ValueStyle.Render(fmt.Sprintf("%.2f WPM", m.Session.CalculateTypingSpeed()))),
		fmt.Sprintf("%s %s", LabelStyle.Render("Raw Speed:"), ValueStyle.Render(fmt.Sprintf("%.2f WPM", m.Session.CalculateRawTypingSpeed()))),
		fmt.Sprintf("%s %s", LabelStyle.Render("Accuracy:"), ValueStyle.Render(fmt.Sprintf("%.2f%%", m.Session.CalculateAccuracy()))),
		formattedTime,
	)

	// Wrap stats in a subtle border or padding
	styledStats := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(CatSurface).
		Padding(1, 3).
		Render(statsBox)

	footer := HelpStyle.Render("[ESC] Quit • [TAB] Restart • [CTRL+P] Command Palette")

	modeSelectorString := ""
	quoteTypeSelectorString := ""

	if m.IsSelectingMode {
		modeSelectorString = getSeletorString(m.CurrentSelector.GetChoices(), m.CurrentSelector.GetCursor())
	}

	if m.IsSelectingQuoteType {
		quoteTypeSelectorString = getSeletorString(m.CurrentSelector.GetChoices(), m.CurrentSelector.GetCursor())
	}

	// Final Assembly
	ui := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"\n",
		styledStats,
		"\n",
		footer,
		"\n",
		modeSelectorString,
		quoteTypeSelectorString,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, ui)
}
