package modes

import (
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
)

type TimeMode struct {
	Duration int
	Language database.Language
}

func (t TimeMode) GetTarget() string {
	return typing.GetDictionnary(t.Language)
}

func (t TimeMode) GetInitialTime() int {
	return t.Duration
}
