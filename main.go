package main

import (
	"github.com/LulianoM/GamesSimulatorAPI/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
