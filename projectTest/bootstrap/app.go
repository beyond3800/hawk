package bootstrap

import (
    _ "projectTest/database/migrations"

    "projectTest/routes"

    "github.com/beyond3800/hawk/engine"
)

func App() *hawk.Hawk {

    app := routes.SetupRoutes()

    return app
}