package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) LobbyView() string {
	logo := CouikASCII3
	if m.CustomDashboard != "" {
		logo = m.CustomDashboard
	}

	header := HeaderStyle.Render(logo)

	subtitle := lipgloss.NewStyle().
		Foreground(CatOverlay).
		SetString("multiplayer lobby").
		String()

	var playersData []string

	if len(m.Players) == 1 {
		playersData = append(playersData, lipgloss.NewStyle().Foreground(CatOverlay).Italic(true).Render("waiting for players..."))
	} else {
		for name := range m.Players {
			var playerLine string
			if name == m.PlayerName {
				dot := lipgloss.NewStyle().Foreground(CatGreen).Render("●")
				nameText := lipgloss.NewStyle().Foreground(CatText).Bold(true).Render(fmt.Sprintf(" %s", name))
				youTag := lipgloss.NewStyle().Foreground(CatSubtext).Render(" (you)")
				playerLine = lipgloss.JoinHorizontal(lipgloss.Left, dot, nameText, youTag)
			} else {
				dot := lipgloss.NewStyle().Foreground(CatOverlay).Render("○")
				nameText := lipgloss.NewStyle().Foreground(CatText).Render(fmt.Sprintf(" %s", name))
				playerLine = lipgloss.JoinHorizontal(lipgloss.Left, dot, nameText)
			}
			playersData = append(playersData, playerLine)
		}
	}

	playerListContent := lipgloss.JoinVertical(lipgloss.Left, playersData...)

	playerCount := lipgloss.NewStyle().
		Foreground(CatSubtext).
		Render(fmt.Sprintf("%d connected", len(m.Players)))

	var statusKey, statusDesc string

	if m.IsHost {
		statusKey = "ctrl+j"
		statusDesc = "start game"
	} else {
		statusKey = ""
		statusDesc = "waiting for host..."
	}

	keyStyle := lipgloss.NewStyle().Foreground(CatLavender).Bold(true)
	descStyle := lipgloss.NewStyle().Foreground(CatOverlay)

	var statusLine string
	if statusKey != "" {
		statusLine = lipgloss.JoinHorizontal(lipgloss.Center,
			keyStyle.Render(statusKey+" "),
			descStyle.Render(statusDesc),
		)
	} else {
		statusLine = descStyle.Render(statusDesc)
	}

	footerKey := lipgloss.NewStyle().Foreground(CatLavender).Bold(true)
	footerDesc := lipgloss.NewStyle().Foreground(CatOverlay)

	footer := lipgloss.JoinHorizontal(lipgloss.Center,
		footerKey.Render("esc "), footerDesc.Render("quit"),
	)

	content := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"\n",
		subtitle,
		"\n",
		playerCount,
		"\n\n",
		playerListContent,
		"\n\n\n",
		statusLine,
		"\n",
		footer,
	)

	return lipgloss.Place(m.TerminalWidth, m.TerminalHeight, lipgloss.Center, lipgloss.Center, content)
}
