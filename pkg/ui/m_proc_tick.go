package ui

import (
	"github.com/fadilix/couik/internal/engine"
	"github.com/fadilix/couik/pkg/ui/core"
)

// These functions are for the implementation
// of the ProcessTick interface
func (m *Model) GetTimeLeft() int {
	return m.TimeLeft
}

func (m *Model) SetTimeLeft(t int) {
	m.TimeLeft = t
}

func (m *Model) Deactivate() {
	m.Active = false
}

func (m *Model) SetState(s core.SessionState) {
	m.State = s
}

func (m *Model) GetSession() *engine.Session {
	return m.Session
}

func (m *Model) IsActive() bool {
	return m.Active
}

func (m *Model) GetTerminalWidth() int {
	return m.TerminalWidth
}

func (m *Model) CacheChart() {
	width := min(max(m.TerminalWidth/3, 20), 40)
	m.CachedChart = DisplayChart(m.Session.WpmSamples, m.Session.TimesSample, width, 10)
}
