package hawk

import (
	"github.com/beyond3800/hawk/config"
	"github.com/beyond3800/hawk/internal/core/database"
	"github.com/beyond3800/hawk/routes"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Boostrap(){
	database.ConnectDatabase()
	config.ConnectRedis()
}
func (a *App) Run(addr string) {
	// start server
	a.Boostrap()
	r := routes.SetupRoutes()

	r.Run(addr)
}
func (a *App) Test(addr string) {
	// start server without database or redis
	r := routes.SetupRoutes()

	r.Run(addr)
}