package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/vinoMamba/naruto-go/internal/database"
)

func init() {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate database",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	createMigrationCmd := &cobra.Command{
		Use:   "create",
		Short: "create migration",
		Run: func(cmd *cobra.Command, args []string) {
			// migrate create -ext sql -dir config/migrations -seq create_users_table
			if len(args) == 0 {
				log.Fatalln("migration name is required")
			}
			exec.Command("migrate", "create", "-ext", "sql", "-dir", "config/migrations", "-seq", args[0]).Run()
		},
	}
	migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "migrate up",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrateUp()
		},
	}
	migrateDownCmd := &cobra.Command{
		Use:   "down",
		Short: "migrate down",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrateDown()
		},
	}
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(createMigrationCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	rootCmd.AddCommand(migrateCmd)
}
