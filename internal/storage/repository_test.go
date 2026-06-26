package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/fadilix/couik/database"
)

func TestJSONRepository_SaveAndGet(t *testing.T) {
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)

	repo := &JSONRepository{}

	testResult := database.TestResult{
		Quote:    "Test quote",
		WPM:      100,
		RawWPM:   110,
		Acc:      95.5,
		Duration: time.Second * 30,
		Date:     time.Now(),
	}

	repo.Save(testResult)

	history, err := repo.GetHistory()
	if err != nil {
		t.Fatalf("expected no err, got %v", err)
	}

	if len(history) != 1 {
		t.Fatalf("expected 1 result, got %d", len(history))
	}

	if history[0].Quote != "Test quote" || history[0].WPM != 100 {
		t.Fatalf("history data mismatch")
	}
}

func TestJSONRepository_CorruptedFile(t *testing.T) {
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)

	repo := &JSONRepository{}

	appDir := filepath.Join(tempHome, ".config", "couik")
	os.MkdirAll(appDir, 0755)
	historyPath := filepath.Join(appDir, "history.json")
	os.WriteFile(historyPath, []byte("INVALID"), 0644)

	testResult := database.TestResult{
		Quote: "Valid test",
		WPM:   50,
	}

	repo.Save(testResult)

	history, err := repo.GetHistory()
	if err != nil {
		t.Fatalf("expected no err, got %v", err)
	}

	if len(history) != 1 {
		t.Fatalf("expected 1 result after reset, got %d", len(history))
	}
}

func TestJSONRepository_RoundTripKeyTimings(t *testing.T) {
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)

	repo := &JSONRepository{}

	timings := []database.KeyTiming{
		{OffsetMs: 120},
		{OffsetMs: 240, Backspace: true},
		{OffsetMs: 360},
	}

	testResult := database.TestResult{
		Quote:      "Round trip",
		WPM:        75,
		KeyTimings: timings,
	}

	if err := repo.Save(testResult); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	history, err := repo.GetHistory()
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if len(history) != 1 {
		t.Fatalf("expected 1 result, got %d", len(history))
	}

	got := history[0].KeyTimings
	if len(got) != len(timings) {
		t.Fatalf("expected %d timings, got %d", len(timings), len(got))
	}
	for i := range timings {
		if got[i] != timings[i] {
			t.Errorf("timing[%d] mismatch: want %+v got %+v", i, timings[i], got[i])
		}
	}
}
