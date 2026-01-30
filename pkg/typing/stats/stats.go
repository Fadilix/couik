package stats

import "time"

func CalculateTypingSpeed(target string, duration time.Duration) float64 {
	wpm := (float64(len(target)) * (60 / duration.Seconds())) / 5
	return wpm
}
