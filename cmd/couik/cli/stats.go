package cli

import (
	"fmt"
	"os"

	"github.com/fadilix/couik/internal/storage"
	"github.com/spf13/cobra"
)

var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "display your stats",
	Run: func(cmd *cobra.Command, args []string) {
		DisplayStats()
		os.Exit(0)
	},
}

func DisplayStats() {
	fmt.Println("Couik Personal best")
	pb := storage.LoadPRs()

	fmt.Printf("PB WPM: %.2f\n", pb.BestWPM)
	fmt.Printf("Lastest test: %s\n", pb.LastTestDate.Format("02 Jan 2006"))
	fmt.Printf("Streak: %d\n", pb.CurrentStreak)
	fmt.Printf("Total tests: %d\n", pb.TotalTests)
}
