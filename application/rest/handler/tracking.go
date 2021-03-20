package handler

import (
	"net/http"

	"github.com/Lgdev07/gocorreios_server/domain/model"
	"github.com/Lgdev07/gocorreios_server/domain/usecase"
	"github.com/Lgdev07/gocorreios_server/infrastructure/repository"
	"github.com/Lgdev07/gocorreios_server/infrastructure/services"
	"github.com/labstack/echo/v4"
)

type TrackingHandler struct {
	TrackingUseCase usecase.TrackingUseCase
}

func (f *TrackingHandler) Show(c echo.Context) error {
	tracking := &model.Tracking{}

	if err := c.Bind(tracking); err != nil {
		return err
	}

	result, err := f.TrackingUseCase.GetTracking(*tracking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func NewTrackingHandler() *TrackingHandler {
	trackingRepository := repository.GoCorreiosRepository{
		Repo: &services.GoCorreiosService{},
	}
	trackingUseCase := usecase.TrackingUseCase{
		TrackingRepository: &trackingRepository,
	}
	trackingHandler := &TrackingHandler{
		TrackingUseCase: trackingUseCase,
	}

	return trackingHandler
}
