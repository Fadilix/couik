package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Base Palette
	CatMauve    = lipgloss.Color("#cba6f7") // Primary/Logo
	CatLavender = lipgloss.Color("#b4befe") // Secondary
	CatSapphire = lipgloss.Color("#74c7ec") // Accent
	CatText     = lipgloss.Color("#cdd6f4") // Standard Text
	CatSubtext  = lipgloss.Color("#a6adc8") // Faint/Muted

	// Functional Colors
	CatGreen   = lipgloss.Color("#a6e3a1") // Correct
	CatRed     = lipgloss.Color("#f38ba8") // Wrong
	CatYellow  = lipgloss.Color("#f9e2af") // Warning/Highlight
	CatOverlay = lipgloss.Color("#6c7086") // Pending/Placeholder
	CatSurface = lipgloss.Color("#313244") // Background highlights
)

var (
	// The characters you've typed correctly (Soothing Green)
	CorrectStyle = lipgloss.NewStyle().Foreground(CatGreen)

	// Incorrect characters (Soft Red)
	WrongStyle = lipgloss.NewStyle().Foreground(CatRed)

	// If the user misses a space, we highlight the background so it's visible
	SpaceStyle = lipgloss.NewStyle().Background(CatRed).Foreground(CatSurface)

	// Characters remaining in the quote (Muted/Faint)
	PendingStyle = lipgloss.NewStyle().Foreground(CatOverlay)

	// The current character under the cursor (Bold & Underlined)
	HighlightStyle = lipgloss.NewStyle().
			Foreground(CatYellow).
			Underline(true).
			Bold(true)

	// Live stats during typing (minimal, no border)
	StatsStyle = lipgloss.NewStyle().
			Foreground(CatOverlay).
			Padding(0, 1)

	HeaderStyle = lipgloss.NewStyle().Foreground(CatMauve)

	// Mode selector styles
	ModeActiveStyle = lipgloss.NewStyle().
			Foreground(CatMauve).
			Bold(true).
			Padding(0, 1).
			MarginRight(1)

	ModeInactiveStyle = lipgloss.NewStyle().
				Foreground(CatOverlay).
				Padding(0, 1).
				MarginRight(1)

	ModeSelectorContainerStyle = lipgloss.NewStyle().
					Padding(0, 1).
					MarginTop(1)

	// for the views
	ViewHeaderStyle = lipgloss.NewStyle().Foreground(CatMauve).Bold(true)

	// Individual stat styling (command palette & config)
	LabelStyle = lipgloss.NewStyle().Foreground(CatLavender).Bold(true).Width(15).Align(lipgloss.Left)
	ValueStyle = lipgloss.NewStyle().Foreground(CatSubtext).Width(30).Align(lipgloss.Right)

	// Footer Section
	HelpStyle = lipgloss.NewStyle().Foreground(CatOverlay).MarginTop(1)
)
