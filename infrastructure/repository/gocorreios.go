package repository

import (
	"strconv"

	"github.com/Lgdev07/gocorreios/fare"
	"github.com/Lgdev07/gocorreios_server/domain/model"
)

type GoCorreiosRepositoryInt interface {
	Fare(fareInterf fare.Interface) ([]byte, error)
	Tracking(cepsList []string) ([]byte, error)
}

type GoCorreiosRepository struct {
	Repo GoCorreiosRepositoryInt
}

func (gc *GoCorreiosRepository) GetFare(fareModel model.Fare) ([]byte, error) {
	newWeight, _ := strconv.ParseFloat(fareModel.Weight, 64)
	newLength, _ := strconv.ParseFloat(fareModel.Lenght, 64)
	newHeight, _ := strconv.ParseFloat(fareModel.Height, 64)
	newWidth, _ := strconv.ParseFloat(fareModel.Width, 64)

	params := fare.Interface{
		Service:        fareModel.Service,
		CepOrigin:      fareModel.CepOrigin,
		CepDestination: fareModel.CepDestination,
		Weight:         newWeight,
		Length:         newLength,
		Height:         newHeight,
		Width:          newWidth,
	}

	result, err := gc.Repo.Fare(params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (gc *GoCorreiosRepository) GetTracking(trackingModel model.Tracking) ([]byte, error) {
	codesList := trackingModel.Codes

	result, err := gc.Repo.Tracking(codesList)
	if err != nil {
		return nil, err
	}

	return result, nil
}
