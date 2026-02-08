package database

type Category int

const (
	Small Category = iota
	Mid
	Thicc
)

// GetQuotes loads quotes the embedded data and returns
// a list of quotes according to the language and category
func GetQuotes(lang Language, category Category) []Quote {
	var maxLength int
	var minLength int

	switch category {
	case Small:
		minLength = 0
		maxLength = 150
	case Mid:
		minLength = 150
		maxLength = 400
	case Thicc:
		minLength = 400
		maxLength = 99999
	}

	quoteData := LoadEmbeddedQuotes(lang)

	var finalQuotes []Quote

	for _, quote := range quoteData.Quotes {
		if quote.Length <= maxLength && quote.Length >= minLength {
			finalQuotes = append(finalQuotes, quote)
		}
	}

	return finalQuotes
}
