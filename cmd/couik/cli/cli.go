package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "couik",
	Short: "Your typing experience brought to the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var (
	History bool
	Help    bool
	Time    int
	Words   int
	File    string
	Text    string
	SetHelp bool
	Lang    string
	Host    int
	Join    string
)

func Init() {
	SetCommand.Flags().BoolVarP(&SetHelp, "help", "h", false, "Show help for the <set> subcommand")

	ConfigCommand.AddCommand(SetCommand)
	RootCmd.AddCommand(ConfigCommand)
	RootCmd.AddCommand(StatsCmd)

	RootCmd.Flags().BoolVarP(&History, "history", "i", false, "Show history")
	RootCmd.Flags().BoolVarP(&Help, "help", "h", false, "Show help")
	RootCmd.Flags().IntVarP(&Time, "time", "t", 0, "Launch timed typing test (seconds: 30, 60, 120)")
	RootCmd.Flags().IntVarP(&Words, "words", "w", 0, "Launch typing test with specified number of words")
	RootCmd.Flags().StringVarP(&File, "file", "f", "", "Launch a typing test with a custom text in a file")
	RootCmd.Flags().StringVarP(&Text, "custom", "c", "", "Launch a typing test with a custom text")
	RootCmd.Flags().StringVarP(&Lang, "lang", "l", "", "Launch a typing test a specific language")
	RootCmd.Flags().IntVarP(&Host, "host", "p", 4217, "Host a game")
	RootCmd.Flags().StringVarP(&Join, "join", "j", "", "Join a game (localip:port)")
}
