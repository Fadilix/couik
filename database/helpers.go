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

	switch category {
	case Small:
		maxLength = 25
	case Mid:
		maxLength = 60
	case Thicc:
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
