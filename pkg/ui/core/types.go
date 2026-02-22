package core

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
)

type SessionState int

const (
	StateTyping SessionState = iota
	StateResults
	StateCommandPalette
	StateConfig
	StateLobby
	StateCountdown
)

// type QuoteType int
//
// const (
// 	Small QuoteType = iota
// 	Mid
// 	Thicc
// )

type ModeConfig struct {
	Target       string
	Source       string
	Duration     int
	InitialWords int
	Language     database.Language
	Category     database.QuoteCategory
	InitialTime  int
}

type TickMsg time.Time

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type TickWpmMsg time.Time

func WPMTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickWpmMsg(t)
	})
}
