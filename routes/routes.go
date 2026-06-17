package routes

import (
	// "log"
	// "time"

	// "github.com/beyond3800/hawk/config"
	// "github.com/beyond3800/hawk/controllers"
	// "github.com/beyond3800/hawk/util"

	// "github.com/beyond3800/hawk/middleware"
	"time"

	"github.com/beyond3800/hawk/controllers"
	"github.com/beyond3800/hawk/core/hawk"
)

// func SetupRoutes() *gin.Engine {
// 	r := gin.Default()
// 	rdb := config.ConnectRedis()
// 	allowedOrigin,err := util.GetFromEnv("ALLOWED_ORIGIN")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	r.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{allowedOrigin},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	}))

// 	userController := controllers.UserController{}
// 	testerController := controllers.TecterController{}

// 	userController.CreateUser()
// 	protected:= r.Group("/protected")
// 	guestRoutes := r.Group("/")
// 	protected.Use(middleware.Auth(rdb))
// 	guestRoutes.POST("test-redis",testerController.SendEmail)
// 	guestRoutes.GET("/",func(c *gin.Context) {
// 		c.JSON(200, "guest Route")
// 	})

// 	return r
// }

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
	userController := controllers.UserController{}
	app.Get("/", func(c *hawk.Context) {
		c.JSON(200, hawk.SuccessResponse{
			Success: "Welcome to Hawk",
			Message: "This is the default route",
			Code: 200,
		})
	})
	app.Post("/users", userController.CreateUser)
	app.Get("/user",userController.Show)
	return app
}
