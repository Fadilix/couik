package database

import (
	"encoding/json"
	"os"
)

// Save saves test result to
// $CONFIG_PATH/couik/history.json
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

// GetHistory retrieves the history of your typing tests
func GetHistory() (History, error) {
	file, err := getHistoryPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var history History

	json.Unmarshal(data, &history)
	return history, nil
}
