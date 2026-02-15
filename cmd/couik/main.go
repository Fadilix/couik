package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/cmd/couik/cli"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/network"
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

	var choosedLanguage database.Language

	if cli.Lang != "" {
		if cli.Lang == "french" {
			choosedLanguage = database.French
		} else {
			choosedLanguage = database.English
		}
	}

	configLang := cli.GetConfig().Language

	if slices.Contains([]string{"french", "english"}, configLang) {
		if configLang == "french" {
			choosedLanguage = database.French
		} else {
			choosedLanguage = database.English
		}
	}

	randomQuote := typing.GetQuoteUseCase(choosedLanguage, database.Mid)

	target := randomQuote.Text

	m := ui.NewModel(target)
	m.CurrentLanguage = choosedLanguage

	if cli.Words > 0 {
		m = m.GetDictionnaryModelWithWords(cli.Words, choosedLanguage)
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

	if cli.Host != 4217 {
		server := network.NewServer()
		go server.Start(strconv.Itoa(cli.Host))

		time.Sleep(100 * time.Millisecond)

		client, err := network.NewClient("localhost:" + strconv.Itoa(cli.Host))
		if err != nil {
			log.Fatal(err)
		}

		m.Multiplayer = true
		m.IsHost = true
		playerName := "Host"
		if cli.Name != "" {
			playerName = cli.Name
		} else {
			fmt.Println("You should provide a name to play multiplayer (add --name to your command)")
			os.Exit(0)
		}
		m.PlayerName = playerName

		err = client.SendJoin(playerName)
		m.Client = client

		if err != nil {
			log.Fatal(err)
		}

	} else if cli.Join != "" {
		client, err := network.NewClient(cli.Join)
		if err != nil {
			log.Fatal(err)
		}

		m.Client = client

		m.Multiplayer = true
		playerName := "Guest"
		if cli.Name != "" {
			playerName = cli.Name
		} else {
			fmt.Println("You should provide a name to play multiplayer (add --name to your command)")
			os.Exit(0)
		}
		m.PlayerName = playerName

		err = client.SendJoin(playerName)
		if err != nil {
			log.Fatal(err)
		}

	}

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
