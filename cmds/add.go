package cmds

import (
	"fmt"
	"strings"

	"github.com/mxrk/cli_task/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, " ")
		err := db.AddTask(arg)
		if err != nil {
			fmt.Println("Failed to add task", err)
			return
		}
	},
}
