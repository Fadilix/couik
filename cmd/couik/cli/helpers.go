package cli

import (
	"fmt"

	"github.com/fadilix/couik/database"
)

func DisplayHistory() {
	fmt.Println("Couik History")
	history, err := database.GetHistory()
	if err != nil {
		fmt.Println(err)
	}

	for i := range history {
		fmt.Printf("Test #%d\n", i+1)
		fmt.Printf("Speed: %.2f\n", history[i].WPM)
		fmt.Printf("Accuracy: %.2f\n", history[i].Acc)
		fmt.Printf("Raw speed: %.2f\n", history[i].RawWPM)
		fmt.Printf("Duration: %.2f seconds\n", history[i].Duration.Seconds())
		fmt.Printf("Quote: %s (%d characters)\n", history[i].Quote, len(history[i].Quote))
		fmt.Printf("Date: %s\n", history[i].Date.Format("02 Jan 2006"))
		fmt.Println()
	}
}

func DisplayHelp() {
	RootCmd.Usage()
}
