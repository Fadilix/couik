package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/pkg/typing/stats"
)

var (
	// Base Palette
	catMauve    = lipgloss.Color("#cba6f7") // Primary/Logo
	catLavender = lipgloss.Color("#b4befe") // Secondary
	catSapphire = lipgloss.Color("#74c7ec") // Accent
	catText     = lipgloss.Color("#cdd6f4") // Standard Text
	catSubtext  = lipgloss.Color("#a6adc8") // Faint/Muted

	// Functional Colors
	catGreen   = lipgloss.Color("#a6e3a1") // Correct
	catRed     = lipgloss.Color("#f38ba8") // Wrong
	catYellow  = lipgloss.Color("#f9e2af") // Warning/Highlight
	catOverlay = lipgloss.Color("#6c7086") // Pending/Placeholder
	catSurface = lipgloss.Color("#313244") // Background highlights
)

var (
	// The characters you've typed correctly (Soothing Green)
	correctStyle = lipgloss.NewStyle().Foreground(catGreen)

	// Incorrect characters (Soft Red)
	wrongStyle = lipgloss.NewStyle().Foreground(catRed)

	// If the user misses a space, we highlight the background so it's visible
	spaceStyle = lipgloss.NewStyle().Background(catRed).Foreground(catSurface)

	// Characters remaining in the quote (Muted/Faint)
	pendingStyle = lipgloss.NewStyle().Foreground(catOverlay)

	// The current character under the cursor (Bold & Underlined)
	highlightStyle = lipgloss.NewStyle().
			Foreground(catYellow).
			Underline(true).
			Bold(true)

	// Your WPM/ACC display (Vibrant Sapphire)
	statsStyle = lipgloss.NewStyle().
			Foreground(catSapphire).
			Bold(true).
			Padding(0, 1).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(catSurface)

	headerStyle = lipgloss.NewStyle().Foreground(catLavender)

	// Mode selector styles
	modeActiveStyle = lipgloss.NewStyle().
			Foreground(catSurface).
			Background(catMauve).
			Bold(true).
			Padding(0, 2).
			MarginRight(1)

	modeInactiveStyle = lipgloss.NewStyle().
				Foreground(catSubtext).
				Background(catSurface).
				Padding(0, 2).
				MarginRight(1)

	modeSelectorContainerStyle = lipgloss.NewStyle().
					Border(lipgloss.RoundedBorder()).
					BorderForeground(catOverlay).
					Padding(0, 1).
					MarginTop(1)
)

func (m Model) View() string {
	if m.State == stateResults || m.timeLeft <= 0 {
		return m.resultsView()
	}

	if m.Quitting {
		return "Closing Couik...\n"
	}

	// Calculate the visible character window (approximately 5 lines worth)
	textWidth := m.TerminalWidth - 50
	if textWidth < 40 {
		textWidth = 40
	}
	charsPerLine := textWidth
	visibleChars := charsPerLine * 5 // ~5 lines of text

	// Calculate the window start position (keep cursor in the middle-ish)
	windowStart := m.Index - (visibleChars / 3)
	if windowStart < 0 {
		windowStart = 0
	}
	windowEnd := windowStart + visibleChars
	if windowEnd > len(m.Target) {
		windowEnd = len(m.Target)
		windowStart = windowEnd - visibleChars
		if windowStart < 0 {
			windowStart = 0
		}
	}

	// Only render the visible portion of text
	var textArea strings.Builder
	for i := windowStart; i < windowEnd; i++ {
		s := string(m.Target[i])
		switch {
		case i < m.Index:
			if m.Results[i] {
				textArea.WriteString(correctStyle.Render(s))
			} else {
				if s == " " {
					textArea.WriteString(spaceStyle.Render(s))
				} else {
					textArea.WriteString(wrongStyle.Render(s))
				}
			}
		case i == m.Index:
			textArea.WriteString(highlightStyle.Render(s))
		default:
			textArea.WriteString(pendingStyle.Render(s))
		}
	}

	var liveWpm, liveAcc float64
	if m.Started && m.Index > 0 {
		correctCount := 0
		for _, r := range m.Results[:m.Index] {
			if r {
				correctCount++
			}
		}
		liveWpm = stats.CalculateTypingSpeed(correctCount, time.Since(m.StartTime))
		liveAcc = stats.CalculateAccuracy(correctCount, m.Index, m.BackSpaceCount)
	}

	wpmDisplay := statsStyle.Render(fmt.Sprintf("WPM: %.2f", liveWpm))
	accDisplay := statsStyle.Render(fmt.Sprintf("ACC: %.2f%%", liveAcc))
	statsRow := lipgloss.JoinHorizontal(lipgloss.Top, wpmDisplay, accDisplay)

	var percent float64
	if m.Mode == timedMode && m.initialTime > 0 {
		percent = float64(m.initialTime-m.timeLeft) / float64(m.initialTime)
	} else {
		percent = float64(m.Index) / float64(len(m.Target))
	}
	bar := m.ProgressBar.ViewAs(percent)

	textAreaStyle := lipgloss.NewStyle().Width(textWidth).Height(5).Align(lipgloss.Left)
	wrappedText := textAreaStyle.Render(textArea.String())

	modeSelectorString := ""

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
	timer := ""
	if m.Mode == timedMode {
		timer = fmt.Sprintf("%d\n", m.timeLeft)
	}

	content := fmt.Sprintf(
		`%s


%s


%s

%s

%s
%s
%s
%s`,
		headerStyle.Render(CouikASCII3),
		wrappedText,
		statsRow,
		bar,
		lipgloss.NewStyle().Faint(true).Render("Press Esc to quit • [SHIFT + TAB] change mode"),
		"\n",
		timer,
		modeSelectorString,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) resultsView() string {
	// Logo Section
	header := lipgloss.NewStyle().Foreground(catMauve).Bold(true).Render(CouikASCII3)

	// Stats Section - Using a box to make it stand out
	statsTitleStyle := lipgloss.NewStyle().Foreground(catSapphire).Bold(true).MarginBottom(1)

	// Individual stat styling
	labelStyle := lipgloss.NewStyle().Foreground(catSubtext).Width(15).Align(lipgloss.Left)
	valueStyle := lipgloss.NewStyle().Foreground(catText).Bold(true)

	// Build the stats block
	statsBox := lipgloss.JoinVertical(lipgloss.Left,
		statsTitleStyle.Render("SESSION PERFORMANCE"),
		fmt.Sprintf("%s %s", labelStyle.Render("Speed:"), valueStyle.Render(fmt.Sprintf("%.2f WPM", m.CalculateTypingSpeed()))),
		fmt.Sprintf("%s %s", labelStyle.Render("Raw Speed:"), valueStyle.Render(fmt.Sprintf("%.2f WPM", m.CalculateRawTypingSpeed()))),
		fmt.Sprintf("%s %s", labelStyle.Render("Accuracy:"), valueStyle.Render(fmt.Sprintf("%.2f%%", m.CalculateAccuracy()))),
	)

	// Wrap stats in a subtle border or padding
	styledStats := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(catSurface).
		Padding(1, 3).
		Render(statsBox)

	// Footer Section
	helpStyle := lipgloss.NewStyle().Foreground(catOverlay).MarginTop(1)
	footer := helpStyle.Render("[TAB] restart • [SHIFT + TAB] change mode • [ESC] quit")

	modeSelectorString := ""

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

	// Final Assembly
	ui := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"\n",
		styledStats,
		"\n",
		footer,
		"\n",
		modeSelectorString,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, ui)
}
