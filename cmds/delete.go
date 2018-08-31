package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete a tasks")
	},
}
