package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var dbCmd = &cobra.Command{
		Use:   "db",
		Short: "db commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("db called")
		},
	}
	rootCmd.AddCommand(dbCmd)
}
