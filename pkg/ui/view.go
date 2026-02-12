package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/pkg/typing/stats"
	"github.com/fadilix/couik/pkg/ui/core"
	"github.com/fadilix/couik/pkg/ui/modes"
)

var dashboardLogo string = CouikASCII3

func (m Model) View() string {
	_, isTimeMode := m.Mode.(*modes.TimeMode)
	if m.State == core.StateResults || (isTimeMode && m.TimeLeft <= 0) {
		return m.resultsView()
	}

	if m.State == core.StateCommandPalette {
		return m.commandPaletteView()
	}

	if m.State == core.StateConfig {
		return m.configView()
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
				textArea.WriteString(CorrectStyle.Render(s))
			} else {
				if s == " " {
					textArea.WriteString(SpaceStyle.Render(s))
				} else {
					textArea.WriteString(WrongStyle.Render(s))
				}
			}
		case i == m.Session.Index:
			textArea.WriteString(HighlightStyle.Render(s))
		default:
			textArea.WriteString(PendingStyle.Render(s))
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

	wpmDisplay := StatsStyle.Render(fmt.Sprintf("WPM: %.2f", liveWpm))
	accDisplay := StatsStyle.Render(fmt.Sprintf("ACC: %.2f%%", liveAcc))
	statsRow := lipgloss.JoinHorizontal(lipgloss.Top, wpmDisplay, accDisplay)

	var percent float64
	if isTimeMode && m.initialTime > 0 {
		percent = float64(m.initialTime-m.TimeLeft) / float64(m.initialTime)
	} else if len(m.Session.Target) > 0 {
		percent = float64(m.Session.Index) / float64(len(m.Session.Target))
	} else {
		percent = 0
	}
	bar := m.ProgressBar.ViewAs(percent)

	textAreaStyle := lipgloss.NewStyle().Width(textWidth).Height(visibleLines).Align(lipgloss.Left)
	wrappedText := textAreaStyle.Render(textArea.String())

	modeSelectorString := ""
	quoteTypeSelectorString := ""

	if m.IsSelectingMode {
		var modeButtons []string
		for i, choice := range m.CurrentSelector.GetChoices() {
			var styledChoice string
			if m.CurrentSelector.GetCursor() == i {
				styledChoice = ModeActiveStyle.Render(choice)
			} else {
				styledChoice = ModeInactiveStyle.Render(choice)
			}
			modeButtons = append(modeButtons, styledChoice)
		}
		buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, modeButtons...)
		modeSelectorString = ModeSelectorContainerStyle.Render(buttonRow)
	}

	if m.IsSelectingQuoteType {
		var quoteTypeButtons []string
		for i, choice := range m.CurrentSelector.GetChoices() {
			var styledChoice string
			if m.CurrentSelector.GetCursor() == i {
				styledChoice = ModeActiveStyle.Render(choice)
			} else {
				styledChoice = ModeInactiveStyle.Render(choice)
			}
			quoteTypeButtons = append(quoteTypeButtons, styledChoice)
		}
		buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, quoteTypeButtons...)
		quoteTypeSelectorString = ModeSelectorContainerStyle.Render(buttonRow)
	}

	timer := ""
	words := ""
	// _, isTimeMode := m.Mode.(*modes.TimeMode)

	if isTimeMode {
		timer = fmt.Sprintf("%d\n", m.TimeLeft)
	} else {
		words = fmt.Sprintf("%d/%d\n", m.Session.Index, len(string(m.Session.Target)))
	}

	renderedLogo := dashboardLogo

	if m.CustomDashboard != "" {
		renderedLogo = m.CustomDashboard
	}

	content := lipgloss.JoinVertical(lipgloss.Center,
		HeaderStyle.Render(renderedLogo),
		"\n",
		wrappedText,
		"\n",
		statsRow,
		bar,
		lipgloss.NewStyle().Faint(true).Render("[Esc] Quit • [CTRL+P] Command Palette"),
		"\n",
		timer,
		words,
		modeSelectorString,
		quoteTypeSelectorString,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, content)
}
