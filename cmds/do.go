package cmds

import (
	"fmt"
	"strings"

	"github.com/mxrk/cli_task/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do a task",
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, " ")
		err := db.DoTask(arg)
		if err != nil {
			fmt.Println("Failed to complete task: ", err)
			return
		}
		fmt.Printf("Marked task number %v as completed.\n", arg)
	},
}
