package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/cmd/couik/cli"
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
		cli.DisplayHistory()
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
