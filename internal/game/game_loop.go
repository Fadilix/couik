package game

import (
	"bufio"
	"fmt"
	"time"

	"github.com/fadilix/couik/pkg/typing/stats"
	"github.com/fatih/color"
)

func GameLoop(target string, r *bufio.Reader) {
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")
	fmt.Print("\rStart typing when you are ready!\n")

	var startTime time.Time
	started := false

	results := make([]bool, len(target))

	for i := 0; i < len(target); {
		fmt.Print("\r")

		for j := 0; j < i; j++ {
			if results[j] {
				color.Set(color.FgGreen)
			} else {
				color.Set(color.FgRed)
			}
			fmt.Print(string(target[j]))
		}

		color.Unset()

		color.Set(color.FgHiBlack)
		fmt.Print(target[i:])
		color.Unset()

		// move the cursor down
		fmt.Print("\033[K")

		if started {
			currentDuration := time.Since(startTime)
			currentCorrect := CountCorrect(results[:i])
			liveWpm := stats.CalculateTypingSpeed(currentCorrect, currentDuration)
			liveAcc := stats.CalculateAccuracy(currentCorrect, len(target[:i]))

			// \033[2B moves cursor down 2 lines
			// \033[2A moves cursor up 2 lines
			fmt.Printf("\033[2B\r\033[KLive WPM : %.2f\n\rLive Acc : %.2f%%\033[3A", liveWpm, liveAcc)
		}

		char, _ := r.ReadByte()
		if char == 127 || char == 8 {
			if i > 0 {
				i--
			}
			continue
		}

		if !started {
			startTime = time.Now()
			started = true
		}

		if char == target[i] {
			results[i] = true
		} else {
			results[i] = false
		}

		i++

		if char == 3 {
			return
		}
	}

	duration := time.Since(startTime)

	correct := CountCorrect(results)
	incorrect := CountIncorrect(results)

	wpm := stats.CalculateTypingSpeed(correct, duration)
	rawWpm := stats.CalculateRawTypingSpeed(correct, incorrect, duration)
	acc := stats.CalculateAccuracy(correct, len(target))

	fmt.Printf("\r\nCongratulations Your typing speed is %.2f WPM\n", wpm)
	fmt.Printf("\rYour raw typing speed is %.2f WPM\n", rawWpm)
	fmt.Printf("\rYour accuracy is %.2f", acc)
}
