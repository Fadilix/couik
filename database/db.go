package database

import (
	"encoding/json"
	"os"
)

func Save(test TestResult) error {
	var history History
	file, err := getHistoryPath()
	if err != nil {
		return err
	}

	if !fileExists(file) {
		initialData := "[]"
		os.WriteFile(file, []byte(initialData), 0o644)
	}

	data, err := os.ReadFile(file)

	if err == nil {
		err := json.Unmarshal(data, &history)
		if err != nil {
			return err
		}
	}

	history = append(history, test)

	newData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, newData, 0o644)
}
