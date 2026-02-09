package database

import (
	_ "embed"
	"encoding/json"
	"log"
)

var (
	//go:embed static/quotes/french.json
	frenchQuotes []byte

	//go:embed static/quotes/english.json
	englishQuotes []byte
)

// Language category
type Language int

const (
	French Language = iota
	English
)

// LoadEmbeddedQuotes loads data from json files
// according to the language (english and french for now)
func LoadEmbeddedQuotes(lang Language) TypingData {
	var quotes TypingData
	var quotesData []byte // from json

	switch lang {
	case French:
		quotesData = frenchQuotes
	default:
		quotesData = englishQuotes
	}

	err := json.Unmarshal(quotesData, &quotes)
	if err != nil {
		log.Fatal("Could not parse embedded JSON", err)
	}
	return quotes
}
