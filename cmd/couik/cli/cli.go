package cli

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "couik",
	Short: "You have to type couiker",
}

var History bool

func Init() {
	RootCmd.Flags().BoolVarP(&History, "history", "i", false, "Show history")
}
