package database

import "time"

type KeyTiming struct {
	OffsetMs  int64 `json:"t"`
	Backspace bool  `json:"b,omitempty"`
}

type TestResult struct {
	RawWPM     float64       `json:"raw_wpm"`
	WPM        float64       `json:"wpm"`
	Acc        float64       `json:"accuracy"`
	Duration   time.Duration `json:"duration"`
	Date       time.Time     `json:"date"`
	Quote      string        `json:"quote"`
	KeyTimings []KeyTiming   `json:"key_timings,omitempty"`
}

type History []TestResult
