package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) resultsView() string {
	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	// Logo Section
	header := lipgloss.NewStyle().Foreground(CatMauve).Bold(true).Render(renderedLogo)

	// Stats Section - Using a box to make it stand out
	StatsTitleStyle := lipgloss.NewStyle().Foreground(CatSapphire).Bold(true).MarginBottom(1)

	// Individual stat styling
	LabelStyle := lipgloss.NewStyle().Foreground(CatSubtext).Width(15).Align(lipgloss.Left)
	ValueStyle := lipgloss.NewStyle().Foreground(CatText).Bold(true)

	// Build the stats block
	statsBox := lipgloss.JoinVertical(lipgloss.Left,
		StatsTitleStyle.Render("SESSION PERFORMANCE"),
		fmt.Sprintf("%s %s", LabelStyle.Render("Speed:"), ValueStyle.Render(fmt.Sprintf("%.2f WPM", m.Session.CalculateTypingSpeed()))),
		fmt.Sprintf("%s %s", LabelStyle.Render("Raw Speed:"), ValueStyle.Render(fmt.Sprintf("%.2f WPM", m.Session.CalculateRawTypingSpeed()))),
		fmt.Sprintf("%s %s", LabelStyle.Render("Accuracy:"), ValueStyle.Render(fmt.Sprintf("%.2f%%", m.Session.CalculateAccuracy()))),
	)

	// Wrap stats in a subtle border or padding
	styledStats := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(CatSurface).
		Padding(1, 3).
		Render(statsBox)

	// Footer Section
	helpStyle := lipgloss.NewStyle().Foreground(CatOverlay).MarginTop(1)
	footer := helpStyle.Render("[TAB] restart • [CTRL + L] restart same • [SHIFT + TAB] change mode • [ESC] quit")

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
