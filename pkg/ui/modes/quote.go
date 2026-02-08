package modes

import (
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
)

type QuoteMode struct {
	Language database.Language
	Category database.Category
}

func (q QuoteMode) GetTarget() string {
	quote := typing.GetQuoteUseCase(q.Language, q.Category)
	return quote.Text
}

func (q QuoteMode) GetInitialTime() int {
	return 0
}
