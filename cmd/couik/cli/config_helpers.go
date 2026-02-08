package cli

import (
	"log"
	"os"
	"slices"
	"strconv"

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
	config := &Config{}

	availableKeys := []string{
		"mode",
		"dashboard_ascii",
		"quote_type",
		"time",
	}

	if !slices.Contains(availableKeys, key) {
		log.Fatal("Can't use this key")
	}

	switch key {
	case "mode":
		if !slices.Contains([]string{"quote", "time", "word"}, value) {
			log.Fatal("Can't use this value")
		}
		config.Mode = value

	case "dashboard_ascii":
		if !database.FileExists(value) {
			log.Fatal("The path is incorrect")
		}
		config.DashboardAscii = value
	case "quote_type":
		if !slices.Contains([]string{"small", "mid", "thicc"}, value) {
			log.Fatal("Can't use this value as quote_type")
		}
		config.QuoteType = value

	case "time":
		if _, err := strconv.Atoi(value); err != nil {
			log.Fatal("You have to use an integer as preferred time")
		}

		val, _ := strconv.Atoi(value)
		config.Time = val
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
