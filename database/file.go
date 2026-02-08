package database

import (
	"errors"
	"os"
	"path/filepath"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}

type fileOption int

const (
	Historyy fileOption = iota
	Config
)

func GetPath(option fileOption) (string, error) {
	var filename string
	if option == Historyy {
		filename = "history.json"
	} else {
		filename = "config.yaml"
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, "couik")
	err = os.MkdirAll(appDir, 0o755)
	if err != nil {
		return "", err
	}

	return filepath.Join(appDir, filename), nil
}
