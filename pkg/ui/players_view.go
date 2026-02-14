package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) PlayersView() string {
	boxStyle := lipgloss.NewStyle().Foreground(CatMauve).Align(lipgloss.Center).Padding(0, 1)

	var s strings.Builder

	myWpm := int(m.Session.CalculateLiveTypingSpeed())
	myXp := fmt.Sprintf("name: %s (YOU) --- speed: %d\n", m.PlayerName, myWpm)
	s.WriteString(myXp)

	for name, opponent := range m.Players {
		if name == m.PlayerName {
			continue
		}
		userStats := fmt.Sprintf("name: %s --- speed: %d\n", name, opponent.WPM)
		s.WriteString(userStats)
	}
	if m.Multiplayer {
		return boxStyle.Render(s.String())
	} else {
		return ""
	}
}
