package cli

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "couik",
	Short: "You have to type couiker",
}

var (
	History bool
	Time    int
	Words   int
	File    string
)

func Init() {
	RootCmd.Flags().BoolVarP(&History, "history", "i", false, "Show history")
	RootCmd.Flags().IntVarP(&Time, "time", "t", 0, "Launch timed typing test (seconds: 30, 60, 120)")
	RootCmd.Flags().IntVarP(&Words, "words", "w", 0, "Launch typing test with specified number of words")
	RootCmd.Flags().StringVarP(&File, "file", "f", "", "Launch a typing test with a custom text in a file")
}
