package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/ui/components"
	"github.com/fadilix/couik/pkg/ui/core"
	"github.com/fadilix/couik/pkg/ui/modes"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case core.TickMsg:
		cmd := m.Mode.ProcessTick(&m)
		return m, cmd

	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height

		m.ProgressBar.Width = msg.Width - 20
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "l", "right":
			m.CurrentSelector.Increment()
		case "h", "left":
			m.CurrentSelector.Decrement()

		case "enter":
			if m.IsSelectingMode {
				selected := m.CurrentSelector.Selected()
				if mode := CreateModeFromSelection(selected, m.CurrentLanguage); mode != nil {
					return m.ApplyMode(mode), nil
				}
			} else if m.IsSelectingQuoteType {
				// selected := m.QuoteTypeChoices[m.QuoteTypeCursor]
				selected := m.CurrentSelector.Selected()

				switch selected {
				case "small":
					m.QuoteType = database.Small
					return m.ApplyMode(modes.NewQuoteMode(modes.WithCategoryQ(database.Small), modes.WithLanguageQ(m.CurrentLanguage))), nil
				case "mid":
					m.QuoteType = database.Mid
					return m.ApplyMode(modes.NewQuoteMode(modes.WithCategoryQ(database.Mid), modes.WithLanguageQ(m.CurrentLanguage))), nil
				case "thicc":
					m.QuoteType = database.Thicc
					return m.ApplyMode(modes.NewQuoteMode(modes.WithCategoryQ(database.Thicc), modes.WithLanguageQ(m.CurrentLanguage))), nil
				}
			}
		}

		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Quitting = true
			return m, tea.Quit

		case tea.KeyCtrlE:
			m.CurrentSelector = components.NewQuoteTypeSelector()
			m.IsSelectingQuoteType = !m.IsSelectingQuoteType
			m.IsSelectingMode = false

		case tea.KeyBackspace:
			if m.State == core.StateTyping {
				m.Session.BackSpace()
			}

		case tea.KeyCtrlP:
			m.State = core.StateCommandPalette

		case tea.KeyCtrlG:
			m.State = core.StateConfig

		case tea.KeyShiftTab:
			m.CurrentSelector = components.NewModeSelector()
			m.IsSelectingMode = !m.IsSelectingMode
			m.IsSelectingQuoteType = false

		case tea.KeyTab:
			if m.State == core.StateResults {
				return m.GetModelFromMode(m), nil
				// switch m.Mode {
				// case quoteMode:
				// 	return m.GetQuoteModel(), nil
				// case wordMode:
				// 	return m.GetDictionnaryModelWithWords(m.InitialWords, m.CurrentLanguage), nil
				// default:
				// 	return m.GetDictionnaryModel(m.initialTime), nil
				// }
			}
		case tea.KeyCtrlL:
			if m.State == core.StateResults {
				return m.ApplyMode(m.Mode), nil
			}
		case tea.KeyCtrlR:
			if m.State == core.StateCommandPalette || m.State == core.StateConfig {
				m.State = core.StateTyping
				return m, nil
			}
			return m.ApplyMode(m.Mode), nil

		case tea.KeyRunes, tea.KeySpace:
			if m.IsSelectingMode {
				return m, nil
			}

			if m.State != core.StateTyping {
				return m, nil
			}

			// Track if we need to start the timer
			_, isTimeMode := m.Mode.(*modes.TimeMode)
			var startTimer bool

			if !m.Active && isTimeMode {
				m.Active = true
				startTimer = true
			}

			if m.Session.Index < len(m.Session.Target) {
				m.Session.Type(msg.String())
			}

			if m.Session.IsFinished() {
				m.State = core.StateResults
				result := database.TestResult{
					RawWPM:   m.Session.CalculateRawTypingSpeed(),
					WPM:      m.Session.CalculateTypingSpeed(),
					Acc:      m.Session.CalculateAccuracy(),
					Duration: m.Session.EndTime.Sub(m.Session.StartTime),
					Quote:    string(m.Session.Target),
					Date:     time.Now(),
				}
				m.Repo.Save(result)
			}

			// Start the timer if needed
			if startTimer {
				return m, core.Tick()
			}
		}
	}
	return m, nil
}
