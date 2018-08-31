package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI taskmanager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Show subcommands and tutorial")
		fmt.Println("-h flag")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
