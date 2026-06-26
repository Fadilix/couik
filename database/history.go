package database

import "time"

type TestResult struct {
	RawWPM   float64       `json:"raw_wpm"`
	WPM      float64       `json:"wpm"`
	Acc      float64       `json:"accuracy"`
	Duration time.Duration `json:"duration"`
	Date     time.Time     `json:"date"`
	Quote    string        `json:"quote"`
}

type History []TestResult
