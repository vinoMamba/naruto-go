package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vinoMamba/naruto-go/internal/router"
)

func init() {
	var srvCmd = &cobra.Command{
		Use:   "server",
		Short: "Run Server",
		Run: func(cmd *cobra.Command, args []string) {
			r := router.New()
			r.Run(":3000")
		},
	}
	rootCmd.AddCommand(srvCmd)
}
