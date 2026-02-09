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

func (ms *QuoteTypeSelector) Increment() {
	if ms.Cursor < len(ms.Choices)-1 {
		ms.Cursor++
	}
}

func (ms *QuoteTypeSelector) Decrement() {
	if ms.Cursor > 0 {
		ms.Cursor--
	}
}

func (ms QuoteTypeSelector) Selected() string {
	return ms.Choices[ms.Cursor]
}
