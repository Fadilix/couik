package modes

import (
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
)

type WordMode struct {
	InitialWords int
	Language     database.Language
}

func (w WordMode) GetTarget() string {
	dict := typing.GetDictionnary(w.Language)
	runes := []rune(dict)
	if w.InitialWords > len(runes) {
		return dict
	}
	return string(runes[:w.InitialWords])
}

func (w WordMode) GetInitialTime() int {
	return 0
}
