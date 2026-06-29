package routes

import (
	"time"

	"github.com/beyond3800/hawk/engine"
)

func SetupRoutes() *hawk.Hawk {
	app := hawk.Default()
	app.Use(hawk.Cors(hawk.CorsConfig{
		AllowOrigins:     []string{"http://localhost:2000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	app.Get("/", func(c *hawk.Context) {
		c.JSON(200, hawk.SuccessResponse{
			Success: "Welcome to Hawk",
			Message: "This is the default route",
			Code: 200,
		})
	})
	return app
}