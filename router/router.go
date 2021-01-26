package router

import (
	"github.com/Lgdev07/gocorreios_server/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/fares", controllers.ShowFare)
}
