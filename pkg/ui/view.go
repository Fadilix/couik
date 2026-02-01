package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/fadilix/couik/pkg/typing/stats"
)

var (
	correctStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	wrongStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	pendingStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	highlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Underline(true)
	statsStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Bold(true).Padding(0, 1)
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
				textArea += wrongStyle.Render(s)
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
		liveAcc = stats.CalculateAccuracy(correctCount, m.Index)
	}

	wpmDisplay := statsStyle.Render(fmt.Sprintf("WPM: %.2f", liveWpm))
	accDisplay := statsStyle.Render(fmt.Sprintf("ACC: %.2f%%", liveAcc))
	statsRow := lipgloss.JoinHorizontal(lipgloss.Top, wpmDisplay, accDisplay)

	percent := float64(m.Index) / float64(len(m.Target))
	bar := m.ProgressBar.ViewAs(percent)

	content := fmt.Sprintf(
		"Couik ↓\n\n%s\n\n%s\n\n%s\n\n%s",
		textArea,
		statsRow,
		bar,
		lipgloss.NewStyle().Faint(true).Render("Press Esc to quit"),
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) resultsView() string {
	s := "✨ Results ✨\n\n"
	s += fmt.Sprintf("WPM:      %.2f\n", m.CalculateTypingSpeed())
	s += fmt.Sprintf("Raw WPM:      %.2f\n", m.CalculateRawTypingSpeed())
	s += fmt.Sprintf("Accuracy: %.2f%%\n", m.CalculateAccuracy())
	s += "\n[Press TAB to restart] • [Press ESC to quit]"

	return s
}
