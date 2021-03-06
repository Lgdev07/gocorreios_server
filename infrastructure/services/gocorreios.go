package services

import (
	"github.com/Lgdev07/gocorreios"
	"github.com/Lgdev07/gocorreios/fare"
)

type GoCorreiosService struct{}

func (gcc *GoCorreiosService) Fare(fareInterf fare.Interface) ([]byte, error) {
	result, err := gocorreios.Fare(fareInterf)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (gcc *GoCorreiosService) Tracking(cepsList []string) ([]byte, error) {
	result, err := gocorreios.Tracking(cepsList)
	if err != nil {
		return nil, err
	}

	return result, nil
}
