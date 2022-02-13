package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/renaldyhidayatt/pencobaan/controllers"
	"github.com/renaldyhidayatt/pencobaan/middleware"
)

func InitialRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", controllers.HelloWorld)

	sumber := api.Group("/sumber")

	sumber.Get("/", controllers.GetSumberAll)
	sumber.Get("/:id", controllers.GetSumberID)
	sumber.Post("/", controllers.CreateSumber)
	sumber.Post("/:id", controllers.UpdateSumber)
	sumber.Delete("/:id", controllers.DeleteSumber)

	uang := api.Group("/uang")

	uang.Get("/", controllers.GetUangAll)
	uang.Get("/:id", controllers.GetUangID)
	uang.Post("/", controllers.CreateUang)
	uang.Post("/:id", controllers.UpdateUang)
	uang.Delete("/:id", controllers.DeleteUang)

	auth := api.Group("/auth")

	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.CreateUsers)

	proctected := api.Group("/proctected")

	proctected.Get("/", middleware.Proctected(), controllers.HelloWorld)

}
