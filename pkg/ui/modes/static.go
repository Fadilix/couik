package modes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
)

type StaticMode struct {
	Target string
}

func (s StaticMode) GetTarget() string {
	return s.Target
}

func NewStaticMode() *StaticMode {
	sm := &StaticMode{
		Target: typing.GetDictionnary(database.English),
	}
	return sm
}

func (s StaticMode) GetInitialTime() int {
	return 0
}

func (s StaticMode) ProcessTick(ctx TickContext) tea.Cmd {
	return nil
}
