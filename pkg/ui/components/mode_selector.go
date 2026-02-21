package components

import "github.com/fadilix/couik/cmd/couik/cli"

type ModeSelector struct {
	Cursor  int
	Choices []string
}

func NewModeSelector() *ModeSelector {
	return &ModeSelector{
		Cursor:  0,
		Choices: cli.TimeWordsVals,
	}
}

func (ms *ModeSelector) Increment() {
	if ms.Cursor < len(ms.Choices)-1 {
		ms.Cursor++
	}
}

func (ms *ModeSelector) Decrement() {
	if ms.Cursor > 0 {
		ms.Cursor--
	}
}

func (ms *ModeSelector) Selected() string {
	return ms.Choices[ms.Cursor]
}

func (ms *ModeSelector) GetChoices() []string {
	return ms.Choices
}

func (ms *ModeSelector) GetCursor() int {
	return ms.Cursor
}
