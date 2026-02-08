package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/pkg/typing/stats"
)

var dashboardLogo string = CouikASCII3

func (m Model) View() string {
	if m.State == stateResults || m.timeLeft <= 0 {
		return m.resultsView()
	}

	if m.Quitting {
		return "Closing Couik...\n"
	}

	// Calculate dimensions
	visibleLines := 3
	textWidth := max(m.TerminalWidth-50, 40)

	var lineStarts []int
	curr := 0
	cursorLine := -1

	for {
		lineStarts = append(lineStarts, curr)

		limit := curr + textWidth
		if limit >= len(m.Session.Target) {
			if cursorLine == -1 && m.Session.Index >= curr {
				cursorLine = len(lineStarts) - 1
			}
			break
		}

		split := limit
		foundSpace := false
		for i := limit; i > curr; i-- {
			if m.Session.Target[i] == ' ' {
				split = i + 1
				foundSpace = true
				break
			}
		}

		if !foundSpace {
			split = limit
		}

		if cursorLine == -1 {
			if m.Session.Index >= curr && m.Session.Index < split {
				cursorLine = len(lineStarts) - 1
			}
		}

		curr = split
		if curr >= len(m.Session.Target) {
			break
		}

		if cursorLine != -1 && len(lineStarts) > cursorLine+3 {
			break
		}
	}

	if cursorLine == -1 {
		cursorLine = len(lineStarts) - 1
	}

	// Calculate window start and end based on lines
	startLineIdx := cursorLine
	startIdx := lineStarts[startLineIdx]

	endIdx := len(m.Session.Target)

	lookaheadLines := 3
	targetLineEndpoint := startLineIdx + lookaheadLines
	if targetLineEndpoint < len(lineStarts) {
		endIdx = lineStarts[targetLineEndpoint]
	}

	windowStart := startIdx
	windowEnd := endIdx
	var textArea strings.Builder
	for i := windowStart; i < windowEnd; i++ {
		s := string(m.Session.Target[i])
		switch {
		case i < m.Session.Index:
			if m.Session.Results[i] {
				textArea.WriteString(correctStyle.Render(s))
			} else {
				if s == " " {
					textArea.WriteString(spaceStyle.Render(s))
				} else {
					textArea.WriteString(wrongStyle.Render(s))
				}
			}
		case i == m.Session.Index:
			textArea.WriteString(highlightStyle.Render(s))
		default:
			textArea.WriteString(pendingStyle.Render(s))
		}
	}

	var liveWpm, liveAcc float64
	if m.Session.Started && m.Session.Index > 0 {
		correctCount := 0
		for _, r := range m.Session.Results[:m.Session.Index] {
			if r {
				correctCount++
			}
		}
		liveWpm = stats.CalculateTypingSpeed(correctCount, time.Since(m.Session.StartTime))
		liveAcc, _ = stats.CalculateAccuracy(correctCount, m.Session.Index, m.Session.BackSpaceCount)
	}

	wpmDisplay := statsStyle.Render(fmt.Sprintf("WPM: %.2f", liveWpm))
	accDisplay := statsStyle.Render(fmt.Sprintf("ACC: %.2f%%", liveAcc))
	statsRow := lipgloss.JoinHorizontal(lipgloss.Top, wpmDisplay, accDisplay)

	var percent float64
	if m.Mode == timedMode && m.initialTime > 0 {
		percent = float64(m.initialTime-m.timeLeft) / float64(m.initialTime)
	} else {
		percent = float64(m.Session.Index) / float64(len(m.Session.Target))
	}
	bar := m.ProgressBar.ViewAs(percent)

	textAreaStyle := lipgloss.NewStyle().Width(textWidth).Height(visibleLines).Align(lipgloss.Left)
	wrappedText := textAreaStyle.Render(textArea.String())

	modeSelectorString := ""
	quoteTypeSelectorString := ""

	if m.IsSelectingMode {
		var modeButtons []string
		for i, choice := range m.Choices {
			var styledChoice string
			if m.Cursor == i {
				styledChoice = modeActiveStyle.Render(choice)
			} else {
				styledChoice = modeInactiveStyle.Render(choice)
			}
			modeButtons = append(modeButtons, styledChoice)
		}
		buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, modeButtons...)
		modeSelectorString = modeSelectorContainerStyle.Render(buttonRow)
	}

	if m.IsSelectingQuoteType {
		var quoteTypeButtons []string
		for i, choice := range m.QuoteTypeChoices {
			var styledChoice string
			if m.QuoteTypeCursor == i {
				styledChoice = modeActiveStyle.Render(choice)
			} else {
				styledChoice = modeInactiveStyle.Render(choice)
			}
			quoteTypeButtons = append(quoteTypeButtons, styledChoice)
		}
		buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, quoteTypeButtons...)
		quoteTypeSelectorString = modeSelectorContainerStyle.Render(buttonRow)
	}

	timer := ""
	if m.Mode == timedMode {
		timer = fmt.Sprintf("%d\n", m.timeLeft)
	}

	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	content := fmt.Sprintf(
		`%s


%s


%s

%s

%s
%s
%s
%s
%s`,
		headerStyle.Render(renderedLogo),
		wrappedText,
		statsRow,
		bar,
		lipgloss.NewStyle().Faint(true).Render("Press Esc to quit • [SHIFT + TAB] change mode"),
		"\n",
		timer,
		modeSelectorString,
		quoteTypeSelectorString,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) resultsView() string {
	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	// Logo Section
	header := lipgloss.NewStyle().Foreground(catMauve).Bold(true).Render(renderedLogo)

	// Stats Section - Using a box to make it stand out
	statsTitleStyle := lipgloss.NewStyle().Foreground(catSapphire).Bold(true).MarginBottom(1)

	// Individual stat styling
	labelStyle := lipgloss.NewStyle().Foreground(catSubtext).Width(15).Align(lipgloss.Left)
	valueStyle := lipgloss.NewStyle().Foreground(catText).Bold(true)

	// Build the stats block
	statsBox := lipgloss.JoinVertical(lipgloss.Left,
		statsTitleStyle.Render("SESSION PERFORMANCE"),
		fmt.Sprintf("%s %s", labelStyle.Render("Speed:"), valueStyle.Render(fmt.Sprintf("%.2f WPM", m.Session.CalculateTypingSpeed()))),
		fmt.Sprintf("%s %s", labelStyle.Render("Raw Speed:"), valueStyle.Render(fmt.Sprintf("%.2f WPM", m.Session.CalculateRawTypingSpeed()))),
		fmt.Sprintf("%s %s", labelStyle.Render("Accuracy:"), valueStyle.Render(fmt.Sprintf("%.2f%%", m.Session.CalculateAccuracy()))),
	)

	// Wrap stats in a subtle border or padding
	styledStats := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(catSurface).
		Padding(1, 3).
		Render(statsBox)

	// Footer Section
	helpStyle := lipgloss.NewStyle().Foreground(catOverlay).MarginTop(1)
	footer := helpStyle.Render("[TAB] restart • [CTRL + L] restart same • [SHIFT + TAB] change mode • [ESC] quit")

	modeSelectorString := ""
	quoteTypeSelectorString := ""

	if m.IsSelectingMode {
		var modeButtons []string
		for i, choice := range m.Choices {
			var styledChoice string
			if m.Cursor == i {
				styledChoice = modeActiveStyle.Render(choice)
			} else {
				styledChoice = modeInactiveStyle.Render(choice)
			}
			modeButtons = append(modeButtons, styledChoice)
		}
		buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, modeButtons...)
		modeSelectorString = modeSelectorContainerStyle.Render(buttonRow)
	}

	if m.IsSelectingQuoteType {
		var quoteTypeButtons []string
		for i, choice := range m.QuoteTypeChoices {
			var styledChoice string
			if m.QuoteTypeCursor == i {
				styledChoice = modeActiveStyle.Render(choice)
			} else {
				styledChoice = modeInactiveStyle.Render(choice)
			}
			quoteTypeButtons = append(quoteTypeButtons, styledChoice)
		}
		buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, quoteTypeButtons...)
		quoteTypeSelectorString = modeSelectorContainerStyle.Render(buttonRow)
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
