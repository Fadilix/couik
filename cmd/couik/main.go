package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/cmd/couik/cli"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui"
)

func main() {
	cli.Init()
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cli.History {
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
		return
	}

	target := typing.GetRandomQuote()
	m := ui.NewModel(target)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
