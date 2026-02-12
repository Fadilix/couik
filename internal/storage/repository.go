package storage

import (
	"encoding/json"
	"os"

	"github.com/fadilix/couik/database"
)

type HistoryRepository interface {
	Save(result database.TestResult) error
	GetHistory() (database.History, error)
}

type JSONRepository struct{}

// Save saves test result to
// $CONFIG_PATH/couik/history.json
func (r *JSONRepository) Save(test database.TestResult) error {
	UpdateStats(test)
	var history database.History
	file, err := database.GetPath(database.Historyy)
	if err != nil {
		return err
	}

	if !database.FileExists(file) {
		initialData := "[]"
		return os.WriteFile(file, []byte(initialData), 0o644)
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

func (r *JSONRepository) GetHistory() (database.History, error) {
	file, err := database.GetPath(database.Historyy)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var history database.History

	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	return history, nil
}
