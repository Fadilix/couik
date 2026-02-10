package ui

import "fmt"

func FormatTime(s int) string {
	hours := s / 3600
	minutes := (s % 3600) / 60
	seconds := s % 60
	return fmt.Sprintf("%02dh %02dmin %02ds", hours, minutes, seconds)
}

// func main() {
// 	fmt.Println(FormatTime(1202470))
// }
