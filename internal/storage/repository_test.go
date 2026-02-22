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
