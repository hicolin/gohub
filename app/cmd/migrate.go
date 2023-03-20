package cmd

import (
	"github.com/spf13/cobra"
	"gohub/database/migrations"
	"gohub/pkg/migrate"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateReset,
		CmdMigrateRefresh,
		CmdMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	migrations.Initialize()
	return migrate.NewMigrator()
}

var CmdMigrateRollback = &cobra.Command{
	Use:     "down",
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runSet,
}

func runSet(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migration",
	Run:   runRefresh,
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}

var CmdMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run all migrations",
	Run:   runFresh,
}

func runFresh(cmd *cobra.Command, args []string) {
	migrator().Fresh()
}
