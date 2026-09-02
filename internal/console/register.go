package console

import (
	"github.com/beyond3800/hawk/internal/console/commands"
	"github.com/beyond3800/hawk/internal/console/commands/db"
	"github.com/beyond3800/hawk/internal/console/commands/migration"
	Make "github.com/beyond3800/hawk/internal/console/commands/make"
)


func registerCommands() {
	// root cli commands
	rootCmd.AddCommand(
		commands.ServeCommand(),
		commands.MigrateCommand(),
		commands.MakeCommands(),
		commands.NewProjectCommand(),
		commands.AuthCommand(),
		commands.VersionCommand(),
		commands.StorageCommand(),
		commands.DbCommand(),
		commands.AppCommand(),
		commands.MigrationCommands(),
	)

	// make cli commands
	commands.MakeCommands().AddCommand(
		Make.AllCommand(),
		Make.ControllerCommand(),
		Make.MiddlewareCommand(),
		Make.MigrationCommand(),
		Make.ModelCommand(),
		Make.RepositoryCommand(),
		Make.ServiceCommand(),
		Make.ResourceCommand(),
		Make.FactoryCommand(),
		Make.SeederCommand(),
		Make.Command(),
		Make.RequestCommand(),
	)

	// migration cli commands
	commands.MigrationCommands().AddCommand(
		migration.RollbackCommand(),
		migration.StatusCommand(),
		migration.FreshCommand(),
		migration.ResetCommand(),
		migration.ReFreshCommand(),
	)

	// db cli commands
	commands.DbCommand().AddCommand(
		db.SeedCommand(),
	)
}


func init() {
	registerCommands()
}