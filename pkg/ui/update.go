package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/ui/components"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		if m.Mode == timedMode {
			m.timeLeft--
			// Check if time is up AFTER decrementing to avoid delay at the end
			if m.timeLeft <= 0 {
				m.Active = false
				m.State = stateResults
				m.Session.EndTime = time.Now()
				return m, nil
			}
			return m, Tick()
		}

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
				switch selected {
				case "15s":
					return m.GetDictionnaryModel(15), nil
				case "30s":
					return m.GetDictionnaryModel(30), nil
				case "60s":
					return m.GetDictionnaryModel(60), nil
				case "120s":
					return m.GetDictionnaryModel(120), nil
				case "quote":
					return m.GetQuoteModel(), nil
				case "words 10":
					return m.GetDictionnaryModelWithWords(10), nil
				case "words 25":
					return m.GetDictionnaryModelWithWords(25), nil
				}
			} else if m.IsSelectingQuoteType {
				// selected := m.QuoteTypeChoices[m.QuoteTypeCursor]
				selected := m.CurrentSelector.Selected()

				switch selected {
				case "small":
					m.QuoteType = small
					return m.GetModelWithQuoteType("small"), nil
				case "mid":
					m.QuoteType = mid
					return m.GetModelWithQuoteType("mid"), nil
				case "thicc":
					m.QuoteType = thicc
					return m.GetModelWithQuoteType("thicc"), nil
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
			if m.State == stateTyping {
				m.Session.BackSpace()
			}

		case tea.KeyCtrlP:
			m.State = stateCommandPalette

		case tea.KeyShiftTab:
			m.CurrentSelector = components.NewModeSelector()
			m.IsSelectingMode = !m.IsSelectingMode
			m.IsSelectingQuoteType = false

		case tea.KeyTab:
			if m.State == stateResults {
				switch m.Mode {
				case quoteMode:
					return m.GetQuoteModel(), nil
				case wordMode:
					return m.GetDictionnaryModelWithWords(m.InitialWords), nil
				default:
					return m.GetDictionnaryModel(m.initialTime), nil
				}
			}
		case tea.KeyCtrlL:
			if m.State == stateResults {
				if m.Mode != timedMode {
					return m.GetModelWithCustomTarget(string(m.Session.Target)), nil
				}
				return m.GetTimeModelWithCustomTarget(m.initialTime, string(m.Session.Target)), nil
			}
		case tea.KeyCtrlR:
			if m.State == stateCommandPalette {
				m.State = stateTyping
				return m, nil
			}
			switch m.Mode {
			case quoteMode:
				var option string
				switch m.QuoteType {
				case mid:
					option = "mid"
				case thicc:
					option = "thicc"
				default:
					option = "small"
				}
				return m.GetModelWithQuoteType(option), nil
			case timedMode:
				return m.GetDictionnaryModel(m.initialTime), nil
			default:
				return m.GetDictionnaryModelWithWords(m.InitialWords), nil
			}

		case tea.KeyRunes, tea.KeySpace:
			if m.IsSelectingMode {
				return m, nil
			}

			if m.State != stateTyping {
				return m, nil
			}

			// Track if we need to start the timer
			var startTimer bool
			if !m.Active && m.Mode == timedMode {
				m.Active = true
				startTimer = true
			}

			if m.Session.Index < len(m.Session.Target) {
				m.Session.Type(msg.String())
			}

			if m.Session.IsFinished() {
				m.State = stateResults
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
				return m, Tick()
			}
		}
	}
	return m, nil
}
