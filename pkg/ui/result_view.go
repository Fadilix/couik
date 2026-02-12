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

	header := ViewHeaderStyle.Render(renderedLogo)

	_, isTimeMode := m.Mode.(*modes.TimeMode)

	wpmVal := fmt.Sprintf("%.0f", m.Session.CalculateTypingSpeed())
	rawVal := fmt.Sprintf("%.0f", m.Session.CalculateRawTypingSpeed())
	accVal := fmt.Sprintf("%.1f%%", m.Session.CalculateAccuracy())

	heroValue := lipgloss.NewStyle().
		Foreground(CatMauve).
		Bold(true).
		Render(wpmVal)

	heroLabel := lipgloss.NewStyle().
		Foreground(CatSubtext).
		Render("wpm")

	heroCard := lipgloss.NewStyle().
		Padding(1, 4).
		Align(lipgloss.Center).
		Render(lipgloss.JoinVertical(lipgloss.Center, heroValue, heroLabel))

	statLabel := lipgloss.NewStyle().Foreground(CatOverlay)
	statValue := lipgloss.NewStyle().Foreground(CatText).Bold(true)
	statDivider := lipgloss.NewStyle().Foreground(CatSurface).Render("  │  ")

	secondaryStats := lipgloss.JoinHorizontal(lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, statValue.Render(rawVal), statLabel.Render("raw")),
		statDivider,
		lipgloss.JoinVertical(lipgloss.Center, statValue.Render(accVal), statLabel.Render("acc")),
	)

	if !isTimeMode {
		durSeconds := int(m.Session.EndTime.Sub(m.Session.StartTime).Seconds())
		durVal := fmt.Sprintf("%ds", durSeconds)
		secondaryStats = lipgloss.JoinHorizontal(lipgloss.Center,
			secondaryStats,
			statDivider,
			lipgloss.JoinVertical(lipgloss.Center, statValue.Render(durVal), statLabel.Render("time")),
		)
	}

	chartWidth := 40
	chartHeight := 12

	chartView := m.CachedChart
	if chartView == "" {
		chartView = DisplayChart(m.Session.WpmSamples, m.Session.TimesSample, chartWidth, chartHeight)
	}

	chartTitle := lipgloss.NewStyle().
		Foreground(CatSubtext).
		Render("wpm over time")

	styledChart := lipgloss.JoinVertical(lipgloss.Center,
		chartTitle,
		"",
		chartView,
	)

	chartContainer := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(CatSurface).
		Padding(1, 2).
		Render(styledChart)

	// ── left column: header + stats ──
	leftColumn := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"",
		heroCard,
		secondaryStats,
	)

	// align left column height to chart
	chartH := lipgloss.Height(chartContainer)
	leftColumn = lipgloss.NewStyle().
		Height(chartH).
		Align(lipgloss.Center, lipgloss.Center).
		Render(leftColumn)

	// ── main row: left + chart ──
	gap := lipgloss.NewStyle().Width(4).Render("")
	mainRow := lipgloss.JoinHorizontal(lipgloss.Center, leftColumn, gap, chartContainer)

	// ── footer ──
	footerKey := lipgloss.NewStyle().Foreground(CatLavender).Bold(true)
	footerDesc := lipgloss.NewStyle().Foreground(CatOverlay)

	footer := lipgloss.JoinHorizontal(lipgloss.Center,
		footerKey.Render("esc "), footerDesc.Render("quit"),
		footerDesc.Render("  •  "),
		footerKey.Render("tab "), footerDesc.Render("restart"),
		footerDesc.Render("  •  "),
		footerKey.Render("ctrl+p "), footerDesc.Render("commands"),
	)

	// ── selectors ──
	modeSelectorString := ""
	quoteTypeSelectorString := ""

	if m.IsSelectingMode {
		modeSelectorString = getSeletorString(m.CurrentSelector.GetChoices(), m.CurrentSelector.GetCursor())
	}

	if m.IsSelectingQuoteType {
		quoteTypeSelectorString = getSeletorString(m.CurrentSelector.GetChoices(), m.CurrentSelector.GetCursor())
	}

	// ── assemble ──
	ui := lipgloss.JoinVertical(lipgloss.Center,
		"",
		mainRow,
		"",
		footer,
		modeSelectorString,
		quoteTypeSelectorString,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, ui)
}
