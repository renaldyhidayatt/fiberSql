package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/renaldyhidayatt/pencobaan/config"
	"github.com/renaldyhidayatt/pencobaan/routes"
)

func main() {
	app := fiber.New()

	config.InitialDatabase()

	routes.InitialRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
