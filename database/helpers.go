package database

type Category int

const (
	small Category = iota
	mid
	thicc
)

func getQuotes(lang Language, category Category) []Quote {
	var maxLength int

	switch category {
	case small:
		maxLength = 25
	case mid:
		maxLength = 60
	case thicc:
		maxLength = 9999
	}

	quoteData := LoadEmbeddedQuotes(lang)

	var finalQuotes []Quote

	for _, quote := range quoteData.Quotes {
		if quote.Length <= maxLength {
			finalQuotes = append(finalQuotes, quote)
		}
	}

	return finalQuotes
}
