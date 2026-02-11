package core

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type SessionState int

const (
	StateTyping SessionState = iota
	StateResults
	StateCommandPalette
	StateConfig
)

type TickMsg time.Time

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
