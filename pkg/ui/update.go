package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.Mode == timedMode {
			m.timeLeft--
			// Check if time is up AFTER decrementing to avoid delay at the end
			if m.timeLeft <= 0 {
				m.Active = false
				m.State = stateResults
				m.EndTime = time.Now()
				return m, nil
			}
			return m, tick()
		}

	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height

		m.ProgressBar.Width = msg.Width - 20
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "l", "right":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "h", "left":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "enter":
			selected := m.Choices[m.Cursor]
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
		}

		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Quitting = true
			return m, tea.Quit

		case tea.KeyBackspace:
			if m.Index > 0 {
				m.Index--
				if m.IsError {
					m.BackSpaceCount++
				}
			}
		case tea.KeyShiftTab:
			m.IsSelectingMode = !m.IsSelectingMode

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
					return m.GetModelWithCustomTarget(m.Target), nil
				}
				return m.GetTimeModelWithCustomTarget(m.initialTime, m.Target), nil
			}
		case tea.KeyCtrlR:
			if m.Mode == quoteMode {
				return m.GetQuoteModel(), nil
			} else {
				return m.GetDictionnaryModel(m.initialTime), nil
			}

		case tea.KeyRunes, tea.KeySpace:
			if m.IsSelectingMode {
				return m, nil
			}

			if m.State == stateResults {
				return m, nil
			}

			// Track if we need to start the timer
			var startTimer bool
			if !m.Active && m.Mode == timedMode {
				m.Active = true
				startTimer = true
			}

			if !m.Started {
				m.StartTime = time.Now()
				m.Started = true
			}

			if m.Index < len(m.Target) {
				typedChar := msg.String()
				if msg.Type == tea.KeySpace {
					typedChar = " "
				}

				isCorrect := typedChar == string(m.Target[m.Index])
				m.IsError = !isCorrect

				m.Results[m.Index] = typedChar == string(m.Target[m.Index])
				m.Index++
			}

			if m.Index >= len(m.Target) {
				m.State = stateResults
				m.EndTime = time.Now()
				// save to database
				result := database.TestResult{
					RawWPM:   m.CalculateRawTypingSpeed(),
					WPM:      m.CalculateTypingSpeed(),
					Acc:      m.CalculateAccuracy(),
					Duration: m.EndTime.Sub(m.StartTime),
					Quote:    m.Target,
					Date:     time.Now(),
				}
				err := database.Save(result)
				if err != nil {
					fmt.Printf("an error occured while trying to save to db : %s\n", err)
				}
			}

			// Start the timer if needed
			if startTimer {
				return m, tick()
			}
		}
	}
	return m, nil
}
