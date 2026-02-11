package cli

import (
	"log"
	"os"
	"slices"
	"strings"

	"github.com/fadilix/couik/database"
	"gopkg.in/yaml.v3"
)

func GetConfig() Config {
	path, err := database.GetPath(database.Config)
	if err != nil {
		log.Fatal("Error while loading your config files")
	}

	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error while reading your config file", err)
	}

	yaml.Unmarshal(data, &config)

	return config
}

func SetConfig(key, value string) {
	config := GetConfig()

	availableKeys := []string{
		"mode",
		"dashboard_ascii",
		"quote_type",
		"time",
		"language",
	}

	if !slices.Contains(availableKeys, key) {
		log.Fatal("Can't use this key")
	}

	switch key {
	case "mode":
		if !slices.Contains([]string{"quote", "time", "words"}, value) {
			log.Fatal("Can't use this value")
		}
		config.Mode = value

	case "dashboard_ascii":
		if !database.FileExists(value) {
			log.Fatal("The path is incorrect")
		}
		config.DashboardASCII = value
	case "quote_type":
		if !slices.Contains([]string{"small", "mid", "thicc"}, value) {
			log.Fatal("Can't use this value as quote_type")
		}
		config.QuoteType = value

	case "time":
		if !slices.Contains([]string{"15s", "30s", "60s", "120s", "quote", "words 10", "words 25"}, value) {
			log.Fatal("You can't use that value as preffered time")
		}
	case "language":
		if !slices.Contains([]string{"french", "english"}, value) {
			log.Fatal("Can't use this language (only english and french available for now)")
		}
		config.Language = value
	}

	path, err := database.GetPath(database.Config)
	if err != nil {
		log.Fatal("Failed to access your config files")
	}

	data, _ := yaml.Marshal(&config)

	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTextFromFile(filepath string) (string, error) {
	quote, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(quote)), nil
}
