package modes

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/core"
)

type TimeMode struct {
	Target      string
	Language    database.Language
	InitialTime int
}

type TimeOption func(*TimeMode)

func NewTimeMode(options ...TimeOption) *TimeMode {
	tm := &TimeMode{
		Target:   typing.GetDictionnary(database.English),
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
	return t.InitialTime
}

func WithLanguageT(language database.Language) TimeOption {
	return func(tm *TimeMode) {
		tm.Language = language
	}
}

func WithTargetT(target string) TimeOption {
	return func(tm *TimeMode) {
		tm.Target = target
	}
}

func WithInitialTimeT(initTime int) TimeOption {
	return func(tm *TimeMode) {
		tm.InitialTime = initTime
	}
}

func (t TimeMode) ProcessTick(ctx TickContext) tea.Cmd {
	if !ctx.IsActive() {
		return nil
	}

	ctx.SetTimeLeft(ctx.GetTimeLeft() - 1)

	if ctx.GetTimeLeft() <= 0 {
		ctx.Deactivate()
		ctx.SetState(core.StateResults)
		ctx.GetSession().EndTime = time.Now()
		ctx.GetSession().Started = false
		ctx.CacheChart()
		return nil
	}

	return core.Tick()
}

func (t TimeMode) GetConfig() core.ModeConfig {
	return core.ModeConfig{
		Target:      t.GetTarget(),
		Language:    t.Language,
		InitialTime: t.InitialTime,
	}
}
