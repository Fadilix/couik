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

	sessionSpeed := m.Session.CalculateTypingSpeed()

	wpmVal := fmt.Sprintf("%.0f", sessionSpeed)
	rawVal := fmt.Sprintf("%.0f", m.Session.CalculateRawTypingSpeed())
	accVal := fmt.Sprintf("%.1f%%", m.Session.CalculateAccuracy())

	heroValue := lipgloss.NewStyle().
		Foreground(CatMauve).
		Bold(true).
		Render(wpmVal)

	heroLabel := lipgloss.NewStyle().
		Foreground(CatSubtext).
		Render("wpm")

	var NewPr string

	if sessionSpeed > m.PRs.BestWPM {
		NewPr = lipgloss.NewStyle().
			Foreground(CatMauve).
			Bold(true).
			Render(" (New PR 󰈸) ")
	}

	heroCard := lipgloss.NewStyle().
		Padding(1, 4).
		Align(lipgloss.Center).
		Render(lipgloss.JoinVertical(lipgloss.Center, heroValue, heroLabel, NewPr))

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

	leftColumn := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"",
		heroCard,
		secondaryStats,
	)

	chartH := lipgloss.Height(chartContainer)
	leftColumn = lipgloss.NewStyle().
		Height(chartH).
		Align(lipgloss.Center, lipgloss.Center).
		Render(leftColumn)

	gap := lipgloss.NewStyle().Width(4).Render("")
	mainRow := lipgloss.JoinHorizontal(lipgloss.Center, leftColumn, gap, chartContainer)

	footerKey := lipgloss.NewStyle().Foreground(CatLavender).Bold(true)
	footerDesc := lipgloss.NewStyle().Foreground(CatOverlay)

	var attribution string
	if m.QuoteSource != "" {
		maxWidth := 80
		label := "Quote from: " + m.QuoteSource
		if len(label) > maxWidth {
			label = label[:maxWidth-3] + "..."
		}
		attribution = lipgloss.NewStyle().Foreground(CatSubtext).Render(label)
	}

	var multiplayerStats string
	if m.Multiplayer {
		multiplayerStats = m.PlayersView()
	}

	footer := lipgloss.JoinHorizontal(lipgloss.Center,
		footerKey.Render("esc "), footerDesc.Render("quit"),
		footerDesc.Render("  •  "),
		footerKey.Render("ctrl+p "), footerDesc.Render("commands"),
	)

	if m.Multiplayer {
		if m.IsHost {
			footer = lipgloss.JoinHorizontal(lipgloss.Center,
				footerKey.Render("esc "), footerDesc.Render("quit"),
				footerDesc.Render("  •  "),
				footerKey.Render("ctrl+j "), footerDesc.Render("restart"),
				footerDesc.Render("  •  "),
				footerKey.Render("ctrl+p "), footerDesc.Render("commands"),
			)
		}
	} else {
		footer = lipgloss.JoinHorizontal(lipgloss.Center,
			footerKey.Render("esc "), footerDesc.Render("quit"),
			footerDesc.Render("  •  "),
			footerKey.Render("tab "), footerDesc.Render("restart"),
			footerDesc.Render("  •  "),
			footerKey.Render("ctrl+p "), footerDesc.Render("commands"),
		)
	}

	modeSelectorString := ""
	quoteTypeSelectorString := ""

	if m.IsSelectingMode {
		modeSelectorString = getSeletorString(m.CurrentSelector.GetChoices(), m.CurrentSelector.GetCursor())
	}

	if m.IsSelectingQuoteType {
		quoteTypeSelectorString = getSeletorString(m.CurrentSelector.GetChoices(), m.CurrentSelector.GetCursor())
	}

	ui := lipgloss.JoinVertical(lipgloss.Center,
		"",
		mainRow,
		"",
		multiplayerStats,
		"",
		attribution,
		"",
		footer,
		modeSelectorString,
		quoteTypeSelectorString,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, ui)
}
