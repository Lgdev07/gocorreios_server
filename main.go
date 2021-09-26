package main

import (
	"os"

	"github.com/Lgdev07/gocorreios_server/application/rest/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://gocorreiosweb-lgdev07.vercel.app"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func InitRoutes(e *echo.Echo) {
	fareHandler := handler.NewFareHandler()
	trackingHandler := handler.NewTrackingHandler()

	e.POST("/fares", fareHandler.Show)
	e.POST("/trackings", trackingHandler.Show)
}
