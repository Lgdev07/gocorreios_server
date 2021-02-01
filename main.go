package main

import (
	"os"

	"github.com/Lgdev07/gocorreios_server/config"
	"github.com/Lgdev07/gocorreios_server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadDotEnv()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(cors.New())
	app.Use(logger.New())

	router.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Listen(":" + port)
}
