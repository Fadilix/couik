package database

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")

	if FileExists(testFile) {
		t.Errorf("expected false")
	}

	os.WriteFile(testFile, []byte("content"), 0644)

	if !FileExists(testFile) {
		t.Errorf("expected true")
	}
}

func TestGetPath(t *testing.T) {
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(tempHome, ".config"))

	tests := []struct {
		option         fileOption
		expectedFile   string
		expectedPreset string
	}{
		{Historyy, "history.json", "[]"},
		{Config, "config.yaml", "{}"},
		{Stats, "stats.json", "[]"},
	}

	for _, tc := range tests {
		t.Run(tc.expectedFile, func(t *testing.T) {
			path, err := GetPath(tc.option)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !FileExists(path) {
				t.Errorf("file was not created")
			}

			content, _ := os.ReadFile(path)
			if string(content) != tc.expectedPreset {
				t.Errorf("expected '%s', got '%s'", tc.expectedPreset, string(content))
			}
		})
	}
}
