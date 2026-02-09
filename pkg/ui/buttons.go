package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func getSeletorString(choices []string, cursor int) string {
	var modeButtons []string
	for i, choice := range choices {
		var styledChoice string
		if cursor == i {
			styledChoice = ModeActiveStyle.Render(choice)
		} else {
			styledChoice = ModeInactiveStyle.Render(choice)
		}
		modeButtons = append(modeButtons, styledChoice)
	}
	buttonRow := lipgloss.JoinHorizontal(lipgloss.Center, modeButtons...)
	return ModeSelectorContainerStyle.Render(buttonRow)
}
