package ui

import (
	"fmt"
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
)

func (m Model) View() string {
	if m.State == stateResults {
		return m.resultsView()
	}

	if m.Quitting {
		return "Closing Couik...\n"
	}

	var textArea string
	for i, char := range m.Target {
		s := string(char)
		switch {
		case i < m.Index:
			if m.Results[i] {
				textArea += correctStyle.Render(s)
			} else {
				if s == " " {
					textArea += spaceStyle.Render(s)
				} else {
					textArea += wrongStyle.Render(s)
				}
			}
		case i == m.Index:
			textArea += highlightStyle.Render(s)
		default:
			textArea += pendingStyle.Render(s)
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

	percent := float64(m.Index) / float64(len(m.Target))
	bar := m.ProgressBar.ViewAs(percent)

	textAreaStyle := lipgloss.NewStyle().Width(m.TerminalWidth - 50).Align(lipgloss.Left)
	wrappedText := textAreaStyle.Render(textArea)

	content := fmt.Sprintf(
		`%s

%s

%s

%s

%s`,
		headerStyle.Render(CouikASCII3),
		wrappedText,
		statsRow,
		bar,
		lipgloss.NewStyle().Faint(true).Render("Press Esc to quit"),
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) resultsView() string {
	// 1. Logo Section
	header := lipgloss.NewStyle().Foreground(catMauve).Bold(true).Render(CouikASCII3)

	// 2. Stats Section - Using a box to make it stand out
	statsTitleStyle := lipgloss.NewStyle().Foreground(catSapphire).Bold(true).MarginBottom(1)

	// Individual stat styling
	labelStyle := lipgloss.NewStyle().Foreground(catSubtext).Width(15).Align(lipgloss.Left)
	valueStyle := lipgloss.NewStyle().Foreground(catText).Bold(true)

	// Build the stats block
	statsBox := lipgloss.JoinVertical(lipgloss.Left,
		statsTitleStyle.Render("📊 SESSION PERFORMANCE"),
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

	// 3. Footer Section
	helpStyle := lipgloss.NewStyle().Foreground(catOverlay).MarginTop(1)
	footer := helpStyle.Render("[TAB] restart  •  [ESC] quit")

	// 4. Final Assembly
	ui := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"\n",
		styledStats,
		"\n",
		footer,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, ui)
}
