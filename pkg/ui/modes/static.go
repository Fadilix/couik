package modes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/core"
)

type StaticMode struct {
	Target string
}

func (qm StaticMode) GetTarget() string {
	return qm.Target
}

func NewStaticMode() *StaticMode {
	sm := &StaticMode{
		Target: typing.GetDictionnary(database.English),
	}
	return sm
}

func (qm StaticMode) GetInitialTime() int {
	return 0
}

func (qm StaticMode) ProcessTick(ctx TickContext) tea.Cmd {
	return nil
}

func (s StaticMode) GetConfig() core.ModeConfig {
	return core.ModeConfig{
		Target: s.Target,
	}
}
