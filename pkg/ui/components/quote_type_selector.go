package components

type QuoteTypeSelector struct {
	Cursor  int
	Choices []string
}

// type Option func(*QuoteTypeSelector)

func NewQuoteTypeSelector() *QuoteTypeSelector {
	selector := &QuoteTypeSelector{
		Cursor:  0,
		Choices: []string{"small", "mid", "thicc"},
	}
	return selector
}

func (qts *QuoteTypeSelector) Increment() {
	if qts.Cursor < len(qts.Choices)-1 {
		qts.Cursor++
	}
}

func (qts *QuoteTypeSelector) Decrement() {
	if qts.Cursor > 0 {
		qts.Cursor--
	}
}

func (qts *QuoteTypeSelector) Selected() string {
	return qts.Choices[qts.Cursor]
}

func (qts *QuoteTypeSelector) GetChoices() []string {
	return qts.Choices
}

func (qts *QuoteTypeSelector) GetCursor() int {
	return qts.Cursor
}
