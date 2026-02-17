package ui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) PlayersView() string {
	if !m.Multiplayer {
		return ""
	}

	type player struct {
		name     string
		wpm      int
		progress float64
		isMe     bool
	}

	var allPlayers []player

	myProgress := 0.0
	if m.Session != nil {
		myProgress = m.Session.Progress()
	}

	allPlayers = append(allPlayers, player{
		name:     m.PlayerName,
		wpm:      int(m.Session.CalculateLiveTypingSpeed()),
		progress: myProgress,
		isMe:     true,
	})

	for name, p := range m.Players {
		if name == m.PlayerName {
			continue
		}
		allPlayers = append(allPlayers, player{
			name:     name,
			wpm:      p.WPM,
			progress: p.Progress,
			isMe:     false,
		})
	}

	sort.Slice(allPlayers, func(i, j int) bool {
		return allPlayers[i].progress > allPlayers[j].progress
	})

	var lines []string
	barWidth := 30

	maxNameLen := 0
	for _, p := range allPlayers {
		l := len(p.name)
		if p.isMe {
			l += 6
		}
		if l > maxNameLen {
			maxNameLen = l
		}
	}
	if maxNameLen > 20 {
		maxNameLen = 20
	}
	if maxNameLen < 10 {
		maxNameLen = 10
	}

	for _, p := range allPlayers {
		nameStr := p.name
		if p.isMe {
			nameStr += " (you)"
		}

		nameStyle := lipgloss.NewStyle().Width(maxNameLen + 2) // +2 padding
		if p.isMe {
			nameStyle = nameStyle.Foreground(CatGreen).Bold(true)
		} else {
			nameStyle = nameStyle.Foreground(CatText)
		}

		wpmStr := fmt.Sprintf("%d wpm", p.wpm)
		wpmStyle := lipgloss.NewStyle().Foreground(CatSubtext).Width(8).Align(lipgloss.Right)

		filledLen := int(p.progress * float64(barWidth))
		if filledLen > barWidth {
			filledLen = barWidth
		}
		if filledLen < 0 {
			filledLen = 0
		}

		var barFilled, barEmpty string
		if filledLen > 0 {
			barFilled = strings.Repeat("━", filledLen)
		}
		if barWidth-filledLen > 0 {
			barEmpty = strings.Repeat("─", barWidth-filledLen)
		}

		barColor := CatOverlay
		if p.isMe {
			barColor = CatMauve
		} else if p.progress >= 1.0 {
			barColor = CatGreen
		}

		bar := lipgloss.NewStyle().Foreground(barColor).Render(barFilled) +
			lipgloss.NewStyle().Foreground(CatSurface).Render(barEmpty)

		row := lipgloss.JoinHorizontal(lipgloss.Center,
			nameStyle.Render(nameStr),
			bar,
			wpmStyle.Render(wpmStr),
		)
		lines = append(lines, row)
	}

	return lipgloss.NewStyle().
		Padding(0, 1).
		MarginTop(1).
		Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}
