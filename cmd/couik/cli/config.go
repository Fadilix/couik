package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "Configure your Couik preferences",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Usage())
		os.Exit(0)
	},
}

var SetCommand = &cobra.Command{
	Use:   "set",
	Short: "Set config preferences",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println(cmd.Usage())
			os.Exit(0)
		}
		key, value := args[0], args[1]

		SetConfig(key, value)
		fmt.Printf("Setting %s to %s", key, value)

		os.Exit(0)
	},
}
