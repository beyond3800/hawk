package commands


func RegisterCommands() {
	// root commands
	RegisterRootCommand(ServeCommand())
	RegisterRootCommand(MakeCommand())
	RegisterRootCommand(RollbackCommand())
	RegisterRootCommand(MigrateCommand())
	RegisterRootCommand(StatusCommand())

	// make commands
	RegisterMakeCommand(RepositoryCommand())
	RegisterMakeCommand(ModelCommand())
	RegisterMakeCommand(ControllerCommand())
	RegisterMakeCommand(ServiceCommand())
	RegisterMakeCommand(AllCommand())
	RegisterMakeCommand(MigrationCommand())
	RegisterMakeCommand(MiddlewareCommand())
}