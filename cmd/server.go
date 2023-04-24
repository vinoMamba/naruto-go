package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var srvCmd = &cobra.Command{
		Use:   "server",
		Short: "Run Server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("server called")
		},
	}
	rootCmd.AddCommand(srvCmd)
}
