package console

import "github.com/beyond3800/hawk/internal/console/commands"


func registerCommands() {
	// root cli commands
	rootCmd.AddCommand(
		commands.ServeCommand(),
		commands.MigrateCommand(),
		commands.RollbackCommand(),
		commands.MakeCommand(),
		commands.StatusCommand(),
		commands.NewProjectCommand(),
	)

	commands.MakeCommand().AddCommand(
		commands.AllCommand(),
		commands.ControllerCommand(),
		commands.MiddlewareCommand(),
		commands.MigrationCommand(),
		commands.ModelCommand(),
		commands.RepositoryCommand(),
		commands.ServiceCommand(),
	)
}


func init() {
	registerCommands()
}