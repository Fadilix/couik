package modes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
)

type QuoteMode struct {
	Language database.Language
	Category database.Category
}

type QuoteOption func(qm *QuoteMode)

func NewQuoteMode(options ...QuoteOption) *QuoteMode {
	qm := &QuoteMode{
		Language: database.English,
		Category: database.Mid,
	}

	for _, option := range options {
		option(qm)
	}

	return qm
}

func (q QuoteMode) GetTarget() string {
	quote := typing.GetQuoteUseCase(q.Language, q.Category)
	return quote.Text
}

func (q QuoteMode) GetInitialTime() int {
	return 0
}

func WithLanguage(language database.Language) QuoteOption {
	return func(qm *QuoteMode) {
		qm.Language = language
	}
}

func WithCategory(category database.Category) QuoteOption {
	return func(qm *QuoteMode) {
		qm.Category = category
	}
}

func (qm QuoteMode) ProcessTick(ctx TickContext) tea.Cmd {
	return nil
}
