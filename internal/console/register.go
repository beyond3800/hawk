package console

import (
	"github.com/beyond3800/hawk/internal/console/commands"
	"github.com/beyond3800/hawk/internal/console/commands/db"
	Make "github.com/beyond3800/hawk/internal/console/commands/make"
)


func registerCommands() {
	// root cli commands
	rootCmd.AddCommand(
		commands.ServeCommand(),
		commands.MigrateCommand(),
		commands.RollbackCommand(),
		commands.MakeCommand(),
		commands.StatusCommand(),
		commands.NewProjectCommand(),
		commands.AuthCommand(),
		commands.VersionCommand(),
		commands.StorageCommand(),
	)

	// make cli commands
	commands.MakeCommand().AddCommand(
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
	)

	// db cli cmmands
	commands.DbCommand().AddCommand(
		db.SeedCommand(),
	)
}


func init() {
	registerCommands()
}