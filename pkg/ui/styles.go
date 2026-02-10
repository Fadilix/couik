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

	// Your WPM/ACC display (Vibrant Sapphire)
	StatsStyle = lipgloss.NewStyle().
			Foreground(CatSapphire).
			Bold(true).
			Padding(0, 1).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(CatSurface)

	HeaderStyle = lipgloss.NewStyle().Foreground(CatLavender)

	// Mode selector styles
	ModeActiveStyle = lipgloss.NewStyle().
			Foreground(CatSurface).
			Background(CatMauve).
			Bold(true).
			Padding(0, 2).
			MarginRight(1)

	ModeInactiveStyle = lipgloss.NewStyle().
				Foreground(CatSubtext).
				Background(CatSurface).
				Padding(0, 2).
				MarginRight(1)

	ModeSelectorContainerStyle = lipgloss.NewStyle().
					Border(lipgloss.RoundedBorder()).
					BorderForeground(CatOverlay).
					Padding(0, 1).
					MarginTop(1)

	// for the views
	// CatMauve
	ViewHeaderStyle = lipgloss.NewStyle().Foreground(CatMauve).Bold(true)

	// Stats Section - Using a box to make it stand out
	StatsTitleStyle = lipgloss.NewStyle().Foreground(CatSapphire).Bold(true).MarginBottom(1)

	// Individual stat styling
	LabelStyle = lipgloss.NewStyle().Foreground(CatSubtext).Width(15).Align(lipgloss.Left)
	ValueStyle = lipgloss.NewStyle().Foreground(CatText).Bold(true)

	// Footer Section
	HelpStyle = lipgloss.NewStyle().Foreground(CatOverlay).MarginTop(1)
)
