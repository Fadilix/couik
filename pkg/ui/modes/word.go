package modes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
)

type WordMode struct {
	InitialWords int
	Language     database.Language
}

type WordOption func(wm *WordMode)

func NewWordMode(options ...WordOption) *WordMode {
	wm := &WordMode{
		InitialWords: 20,
		Language:     database.English,
	}

	for _, option := range options {
		option(wm)
	}

	return wm
}

func WithInitialWords(n int) WordOption {
	return func(wm *WordMode) {
		wm.InitialWords = n
	}
}

// func WithLanguage(language database.Language) WordOption {
// 	return func(wm *WordMode) {
// 		wm.Language = language
// 	}
// }

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

func (w WordMode) ProcessTick(ctx TickContext) tea.Cmd {
	return nil
}
