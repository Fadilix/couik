package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/pkg/ui"
)

func main() {
	target := "The quick brown fox jumps over the lazy dog."
	m := ui.NewModel(target)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
