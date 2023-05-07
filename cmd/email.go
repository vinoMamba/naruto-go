package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vinoMamba/naruto-go/internal/email"
)

func init() {
	var emailCmd = &cobra.Command{
		Use:   "email",
		Short: "send a email",
		Run: func(cmd *cobra.Command, args []string) {
			email.Send()
		},
	}
	rootCmd.AddCommand(emailCmd)
}
