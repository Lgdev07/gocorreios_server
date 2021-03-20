package handler

import (
	"net/http"

	"github.com/Lgdev07/gocorreios_server/domain/model"
	"github.com/Lgdev07/gocorreios_server/domain/usecase"
	"github.com/Lgdev07/gocorreios_server/infrastructure/repository"
	"github.com/Lgdev07/gocorreios_server/infrastructure/services"
	"github.com/labstack/echo/v4"
)

type FareHandler struct {
	FareUseCase usecase.FareUseCase
}

func (f *FareHandler) Show(c echo.Context) error {
	fare := &model.Fare{}

	if err := c.Bind(fare); err != nil {
		return err
	}

	result, err := f.FareUseCase.GetFare(*fare)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func NewFareHandler() *FareHandler {
	fareRepository := repository.GoCorreiosRepository{
		Repo: &services.GoCorreiosService{},
	}
	fareUseCase := usecase.FareUseCase{
		FareRepository: &fareRepository,
	}
	fareHandler := &FareHandler{
		FareUseCase: fareUseCase,
	}

	return fareHandler
}
