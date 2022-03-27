package routes

import (
	"github.com/LulianoM/GamesSimulatorAPI/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	api.Get("healthcheck", controllers.Healthcheck)
	api.Post("jokenpo", controllers.Jokenpo)
}
