package ui

import (
	"github.com/fadilix/couik/database"
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

func (m *Model) SaveResult() {
	if m.Session == nil {
		return
	}

	result := database.TestResult{
		RawWPM:   m.Session.CalculateRawTypingSpeed(),
		WPM:      m.Session.CalculateTypingSpeed(),
		Acc:      m.Session.CalculateAccuracy(),
		Duration: m.Session.EndTime.Sub(m.Session.StartTime),
		Quote:    string(m.Session.Target),
		Date:     m.Session.EndTime,
	}
	m.Repo.Save(result)

	if m.Multiplayer {
		go m.Client.SendUpdate(m.PlayerName, int(result.WPM), 1.0, true)
	}
}
