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
	Stats
)

func GetPath(option fileOption) (string, error) {
	var filename string
	switch option {
	case Historyy:
		filename = "history.json"
	case Config:
		filename = "config.yaml"
	default:
		filename = "stats.json"
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

	fullpath := filepath.Join(appDir, filename)

	if !FileExists(fullpath) {
		err := os.WriteFile(fullpath, []byte(" "), 0o644)
		if err != nil {
			return "", err
		}
	}

	return fullpath, nil
}
