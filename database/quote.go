package database

type TypingData struct {
	Language string  `json:"language"`
	Quotes   []Quote `json:"quotes"`
}

type Quote struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Length int    `json:"length"`
}
