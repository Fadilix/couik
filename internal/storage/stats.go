package storage

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/fadilix/couik/database"
)

// user personal best
type Stats struct {
	BestWPM       float64   `json:"best_wpm"`
	CurrentStreak int       `json:"current_streak"`
	LastTestDate  time.Time `json:"last_test_date"`
	TotalTests    int       `json:"total_tests"`
}

func UpdateStats(test database.TestResult) {
	statsConfig := LoadPRs()

	if test.WPM > statsConfig.BestWPM {
		statsConfig.BestWPM = test.WPM
	}

	ty, tm, td := statsConfig.LastTestDate.Date()
	y, m, d := time.Now().Date()

	if ty == y && tm == m && td != d && td+1 == d {
		statsConfig.CurrentStreak++
	}

	statsConfig.TotalTests++
	statsConfig.LastTestDate = time.Now()

	file, err := database.GetPath(database.Stats)
	if err != nil {
		log.Fatal("Error accessing you stats path")
	}

	newData, err := json.MarshalIndent(statsConfig, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(file, newData, 0o644)
	if err != nil {
		log.Fatal("Error saving your stats")
	}
}

func LoadPRs() Stats {
	var statsConfig Stats
	file, err := database.GetPath(database.Stats)
	if err != nil {
		log.Fatal(err)
	}

	if !database.FileExists(file) {
		initialData := "{}"
		os.WriteFile(file, []byte(initialData), 0o644)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading the stats file")
	}

	err = json.Unmarshal(data, &statsConfig)
	if err != nil {
		// log.Printf("Error retrieving stats data: %v. Resetting stats.", err)
		return Stats{}
	}
	return statsConfig
}
