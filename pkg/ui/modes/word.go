package modes

import "github.com/fadilix/couik/pkg/typing"

type WordMode struct {
	InitialWords int
}

func (w WordMode) GetTarget() string {
	quote := typing.GetDictionnary()[:w.InitialWords]
	return quote
}

func (w WordMode) GetInitialTime() int {
	return 0
}
