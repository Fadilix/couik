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
		cli.DisplayHistory()
		return
	}

	if cli.Help {
		cli.DisplayHelp()
		return
	}

	randomQuote := typing.GetQuoteUseCase(database.English, database.Mid)
	target := randomQuote.Text

	m := ui.NewModel(target)

	if cli.Words > 0 {
		m = m.GetDictionnaryModelWithWords(cli.Words)
	} else if cli.Time > 0 {
		m = m.GetDictionnaryModel(cli.Time)
	} else if cli.File != "" {
		quote, err := typing.GetQuoteFromFile(cli.File)
		if err != nil {
			fmt.Printf("An error occurred while trying to retrieve text in your file %s\n", err)
		}
		target = quote
		m = ui.NewModel(target)
	} else if cli.Text != "" {
		target = cli.Text
		m = ui.NewModel(target)
	}

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
