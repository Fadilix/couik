package ui

import (
	"encoding/json"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fadilix/couik/database"
	"github.com/fadilix/couik/pkg/network"
	"github.com/fadilix/couik/pkg/typing"
	"github.com/fadilix/couik/pkg/ui/components"
	"github.com/fadilix/couik/pkg/ui/core"
	"github.com/fadilix/couik/pkg/ui/modes"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case core.TickMsg:
		if m.State != core.StateTyping {
			return m, nil
		}
		cmd := m.Mode.ProcessTick(&m)
		return m, cmd
	case core.TickWpmMsg:
		if m.State == core.StateCountdown {
			m.Countdown--
			if m.Countdown <= 0 {
				m.State = core.StateTyping
				m.Active = true
				m.Session.Start()
			}
			return m, core.WPMTick()
		}

		if m.State == core.StateTyping && m.Session.Started {
			if m.Multiplayer && !m.Session.IsFinished() {
				go m.Client.SendUpdate(m.PlayerName, int(m.Session.CalculateLiveTypingSpeed()), m.Session.Progress(), false)
			}
			wpm := m.Session.CalculateLiveTypingSpeed()
			if m.State == core.StateTyping {
				m.Session.AddWpmSample(wpm)
				m.Session.AddTimesSample(time.Time(msg))
				return m, core.WPMTick()
			}
		}
		return m, nil

	case network.Message:
		switch msg.Type {
		case network.MsgJoin:
			var joinPayload network.JoinPayload
			if err := json.Unmarshal(msg.Payload, &joinPayload); err != nil {
				log.Println(err)
			} else {
				m.Mu.Lock()
				m.Players[joinPayload.PlayerName] = &network.UpdatePayload{
					PlayerName: joinPayload.PlayerName,
					WPM:        0,
					Progress:   0,
				}
				m.Mu.Unlock()

				if m.IsHost {
					m.Client.SendStart(string(m.Session.Target), 0)
				}
			}

		case network.MsgStart:
			var startLimit network.StartPayload
			if err := json.Unmarshal(msg.Payload, &startLimit); err != nil {
				log.Println(err)
			}

			m.Mu.Lock()
			for _, p := range m.Players {
				p.WPM = 0
				p.Progress = 0
				p.Completed = false
			}
			m.Mu.Unlock()

			m = m.ApplyMode(modes.NewQuoteMode(modes.WithCustomQuote(startLimit.Text)))

			if startLimit.Countdown > 0 {
				m.Countdown = startLimit.Countdown
				m.State = core.StateCountdown
				m.Active = false
				return m, tea.Batch(core.WPMTick(), WaitForNetworkMsg(m.Client))
			}

			return m, WaitForNetworkMsg(m.Client)

		case network.MsgUpdate:
			var updatePayload network.UpdatePayload
			if err := json.Unmarshal(msg.Payload, &updatePayload); err != nil {
				log.Println(err)
			} else {
				m.Mu.Lock()
				m.Players[updatePayload.PlayerName] = &updatePayload
				m.Mu.Unlock()
			}
		}
		return m, WaitForNetworkMsg(m.Client)
	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height

		m.ProgressBar.Width = min(msg.Width-20, 45)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "l", "right":
			m.CurrentSelector.Increment()
		case "h", "left":
			m.CurrentSelector.Decrement()
		case "ctrl+n":
			switch m.CurrentLanguage {
			case database.English:
				m.CurrentLanguage = database.French
			case database.French:
				m.CurrentLanguage = database.English
			}

			cfg := m.Mode.GetConfig()

			if cfg.InitialWords > 0 {
				return m.GetDictionnaryModelWithWords(cfg.InitialWords, m.CurrentLanguage), nil
			}

			if cfg.InitialTime > 0 {
				return m.GetDictionnaryModel(cfg.InitialTime), nil
			}

			return m.ApplyMode(modes.NewQuoteMode(
				modes.WithCategoryQ(cfg.Category),
				modes.WithLanguageQ(m.CurrentLanguage),
				modes.WithTargetQ(typing.GetQuoteUseCase(m.CurrentLanguage, cfg.Category).Text),
			)), nil

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
		case tea.KeyCtrlJ:
			if m.IsHost && m.Multiplayer {
				if m.State == core.StateResults {
					newQuote := typing.GetQuoteUseCase(m.CurrentLanguage, m.QuoteType)
					m.Client.SendStart(newQuote.Text, 5)
				} else {
					m.Client.SendStart(string(m.Session.Target), 5)
				}
			}

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
			if m.State == core.StateResults && !m.Multiplayer {
				return m.GetModelFromMode(m), nil
			}
		case tea.KeyCtrlL:
			if m.State == core.StateResults && !m.Multiplayer {
				return m.ApplyMode(m.Mode, WithSameQuote(string(m.Session.Target))), nil
			}
		case tea.KeyCtrlR:
			// log.Println("i pressed ctrl+r")
			// log.Println(m.State)
			if m.State == core.StateCommandPalette || m.State == core.StateConfig {
				m.State = core.StateTyping
				return m, nil
			}
			return m.GetModelFromMode(m), nil

		case tea.KeyRunes, tea.KeySpace:
			// if m.Ismu

			if m.IsSelectingMode {
				return m, nil
			}

			if m.State != core.StateTyping {
				return m, nil
			}

			// Track if we need to start the timer
			_, isTimeMode := m.Mode.(*modes.TimeMode)
			wasStarted := m.Session.Started

			if !m.Active && isTimeMode {
				m.Active = true
			}

			if m.Session.Index < len(m.Session.Target) {
				m.Session.Type(msg.String())
			}

			if m.Session.IsFinished() {
				m.State = core.StateResults
				m.Session.Started = false
				m.Active = false
				m.CachedChart = DisplayChart(m.Session.WpmSamples, m.Session.TimesSample, min(max(m.TerminalWidth/3, 20), 40), 10)
				result := database.TestResult{
					RawWPM:   m.Session.CalculateRawTypingSpeed(),
					WPM:      m.Session.CalculateTypingSpeed(),
					Acc:      m.Session.CalculateAccuracy(),
					Duration: m.Session.EndTime.Sub(m.Session.StartTime),
					Quote:    string(m.Session.Target),
					Date:     time.Now(),
				}
				m.Repo.Save(result)

				if m.Multiplayer {
					go m.Client.SendUpdate(m.PlayerName, int(m.Session.CalculateTypingSpeed()), 1.0, true)
				}
			}

			// Start the timers if just started
			var cmds []tea.Cmd
			if !wasStarted && m.Session.Started {
				cmds = append(cmds, core.WPMTick())
				if isTimeMode {
					cmds = append(cmds, core.Tick())
				}
			}

			if len(cmds) > 0 {
				return m, tea.Batch(cmds...)
			}
		}
	}
	return m, nil
}
