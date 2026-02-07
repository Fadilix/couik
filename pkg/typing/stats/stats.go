package stats

import "time"

func CalculateTypingSpeed(correctChars int, duration time.Duration) float64 {
	wpm := ((float64(correctChars)) * (60 / duration.Seconds())) / 5
	return wpm
}

func CalculateRawTypingSpeed(correctChars, incorrectChars int, duration time.Duration) float64 {
	wpm := ((float64(correctChars) + float64(incorrectChars)) * (60 / duration.Seconds())) / 5
	return wpm
}

func CalculateAccuracy(correctChars, allChars, backspaceCount int) float64 {
	acc := (float64(correctChars-backspaceCount) / float64(allChars)) * 100
	return acc
}
