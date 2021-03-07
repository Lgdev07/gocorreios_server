package main

import (
	"os"

	"github.com/Lgdev07/gocorreios_server/application/rest/handler"
	"github.com/Lgdev07/gocorreios_server/domain/usecase"
	"github.com/Lgdev07/gocorreios_server/infrastructure/repository"
	"github.com/Lgdev07/gocorreios_server/infrastructure/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	InitRoutes(e)

	e.Use(middleware.CORS())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func InitRoutes(e *echo.Echo) {
	fareRepository := repository.GoCorreiosRepository{
		Repo: &services.GoCorreiosService{},
	}
	fareUseCase := usecase.FareUseCase{
		FareRepository: &fareRepository,
	}
	fareHandler := handler.FareHandler{
		FareUseCase: fareUseCase,
	}

	e.POST("/fares", fareHandler.Show)
}
