package modes

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/core"
)

type TimeMode struct {
	Duration int
	Language database.Language
}

type TimeOption func(tm *TimeMode)

func NewTimeMode(options ...TimeOption) *TimeMode {
	tm := &TimeMode{
		Duration: 60,
		Language: database.English,
	}

	for _, option := range options {
		option(tm)
	}
	return tm
}

func (t TimeMode) GetTarget() string {
	return typing.GetDictionnary(t.Language)
}

func (t TimeMode) GetInitialTime() int {
	return t.Duration
}

func (t TimeMode) ProcessTick(ctx TickContext) tea.Cmd {
	ctx.SetTimeLeft(ctx.GetTimeLeft() - 1)

	if ctx.GetTimeLeft() <= 0 {
		ctx.Deactivate()
		ctx.SetState(core.StateResults)
		ctx.GetSession().EndTime = time.Now()
		return nil
	}

	return core.Tick()
}
