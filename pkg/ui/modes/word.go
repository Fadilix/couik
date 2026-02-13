package modes

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/core"
)

type WordMode struct {
	InitialWords int
	Target       string
	Language     database.Language
}

type WordOption func(wm *WordMode)

func NewWordMode(options ...WordOption) *WordMode {
	wm := &WordMode{
		Target:       typing.GetDictionnary(database.English),
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

func WithLanguageW(language database.Language) WordOption {
	return func(wm *WordMode) {
		wm.Language = language
	}
}

func WithTargetW(target string) WordOption {
	return func(wm *WordMode) {
		wm.Target = target
	}
}

func (w WordMode) GetTarget() string {
	dict := typing.GetDictionnary(w.Language)
	// TODO: come back later to check if the number of words does not exceed

	count := 0
	var res []rune
	for _, char := range dict {
		res = append(res, char)
		if char == ' ' {
			count++
		}
		if count == w.InitialWords {
			break
		}
	}

	return strings.TrimSpace(string(res))
}

func (w WordMode) GetInitialTime() int {
	return 0
}

func (w WordMode) ProcessTick(ctx TickContext) tea.Cmd {
	return nil
}

func (w WordMode) GetConfig() core.ModeConfig {
	return core.ModeConfig{
		Target:       w.GetTarget(),
		InitialWords: w.InitialWords,
		Language:     w.Language,
	}
}
