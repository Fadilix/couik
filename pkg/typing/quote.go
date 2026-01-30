package typing

import (
	"math/rand"
)

var quotes = []string{
	"The only way to do great work is to love what you do.",
	"Talk is cheap. Show me the code.",
	"Success is not final, failure is not fatal: it is the courage to continue that counts.",
	"In the beginning God created the heaven and the earth.",
	"Computers are useless. They can only give you answers.",
	"The quick brown fox jumps over the lazy dog.",
	"To be, or not to be, that is the question.",
	"Stay hungry, stay foolish.",
	"Code is like humor. When you have to explain it, it’s bad.",
	"The best way to predict the future is to invent it.",
}

func GetRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
