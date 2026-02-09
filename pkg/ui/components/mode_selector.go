package components

type ModeSelector struct {
	Cursor  int
	Choices []string
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

func (ms ModeSelector) Selected() string {
	return ms.Choices[ms.Cursor]
}
