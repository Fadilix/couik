package modes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/internal/engine"
	"github.com/fadilix/couik/pkg/ui/core"
)

type ModeStrategy interface {
	GetTarget() string
	GetInitialTime() int
	ProcessTick(TickContext) tea.Cmd
}

type TickContext interface {
	GetTimeLeft() int
	SetTimeLeft(t int)
	Deactivate()
	SetState(s core.SessionState)
	GetSession() *engine.Session
}
