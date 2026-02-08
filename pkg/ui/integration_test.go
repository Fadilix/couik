package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestIntegrationGameLoop(t *testing.T) {
	target := "hello world"
	m := NewModel(target)
	m = m.GetModelWithCustomTarget(target)

	if m.Index != 0 {
		t.Errorf("Expected initial index 0, got %d", m.Index)
	}

	for i, char := range target {
		var msg tea.KeyMsg
		if char == ' ' {
			msg = tea.KeyMsg{Type: tea.KeySpace, Runes: []rune{' '}}
		} else {
			msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{char}}
		}

		// Update the model
		updatedModel, _ := m.Update(msg)
		m = updatedModel.(Model)

		// Verify state after each keystroke
		if m.Index != i+1 {
			t.Errorf("Expected index %d after typing '%c', got %d", i+1, char, m.Index)
		}
		if !m.Results[i] {
			t.Errorf("Expected correct result for char '%c' at index %d", char, i)
		}
	}

	// Get the result
	if m.State != stateResults {
		t.Errorf("Expected state to be stateResults (%d), got %d", stateResults, m.State)
	}

	// Verify accuracy is 100%
	acc := m.CalculateAccuracy()
	if acc != 100.0 {
		t.Errorf("Expected accuracy 100.0, got %f", acc)
	}
}
