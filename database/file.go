package database

import (
	"errors"
	"os"
	"path/filepath"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}

func getHistoryPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, "couik")
	err = os.MkdirAll(appDir, 0o755)
	if err != nil {
		return "", err
	}

	return filepath.Join(appDir, "history.json"), nil
}
