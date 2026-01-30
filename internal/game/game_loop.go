package game

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/fadilix/couik/pkg/typing/stats"
)

func GameLoop(target string, r *bufio.Reader) {
	// when the terminal is in raw mode we use \r to have the same
	// behavior as on Cooked Mode
	fmt.Printf("Type this ↓ \r\n\n %s \r\n\n", target)
	fmt.Print("Type enter when you are ready to type\r\n")

	os.Stdin.Read(make([]byte, 1))

	startTime := time.Now()

	visual := ""
	for i := 0; i < len(target); {
		char, _ := r.ReadByte()

		if char == target[i] {
			fmt.Print(string(char))
			visual += string(char)
			i++
		}

		if char == 3 {
			return
		}
	}
	duration := time.Since(startTime)

	wpm := stats.CalculateTypingSpeed(target, duration)

	fmt.Printf("\r\nYour typing speed is %.2f", wpm)
	fmt.Printf("\r\n\nType any character to quit")
	os.Stdin.Read(make([]byte, 1))
}
