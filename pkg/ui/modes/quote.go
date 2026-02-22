package modes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/core"
)

type QuoteMode struct {
	Target   string
	Source   string
	Language database.Language
	Category database.QuoteCategory
}

type QuoteOption func(qm *QuoteMode)

func NewQuoteMode(options ...QuoteOption) *QuoteMode {
	quote := typing.GetQuoteUseCase(database.English, database.Mid)
	qm := &QuoteMode{
		Target:   quote.Text,
		Source:   quote.Source,
		Language: database.English,
		Category: database.Mid,
	}

	for _, option := range options {
		option(qm)
	}

	return qm
}

func (qm QuoteMode) GetTarget() string {
	if qm.Target != "" {
		return qm.Target
	}
	quote := typing.GetQuoteUseCase(qm.Language, qm.Category)
	return quote.Text
}

func (qm QuoteMode) GetInitialTime() int {
	return 0
}

func WithLanguageQ(language database.Language) QuoteOption {
	return func(qm *QuoteMode) {
		qm.Language = language
	}
}

func WithTargetQ(target string) QuoteOption {
	return func(qm *QuoteMode) {
		qm.Target = target
	}
}

func WithCategoryQ(category database.QuoteCategory) QuoteOption {
	return func(qm *QuoteMode) {
		qm.Category = category
	}
}

func WithCustomQuote(target string) QuoteOption {
	return func(qm *QuoteMode) {
		qm.Target = target
	}
}

func (qm QuoteMode) ProcessTick(ctx TickContext) tea.Cmd {
	return nil
}

func (qm QuoteMode) GetConfig() core.ModeConfig {
	return core.ModeConfig{
		Target:   qm.GetTarget(),
		Source:   qm.Source,
		Language: qm.Language,
		Category: qm.Category,
	}
}
