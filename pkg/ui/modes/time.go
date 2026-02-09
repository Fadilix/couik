package modes

import "github.com/fadilix/couik/pkg/typing"

type TimeMode struct {
	Duration int
}

func (m TimeMode) GetTarget() string {
	return typing.GetDictionnary()
}

func (m TimeMode) GetInitialTime() int {
	return m.Duration
}
