package usecase

import (
	"encoding/json"

	"github.com/Lgdev07/gocorreios_server/domain/model"
)

type FareUseCase struct {
	FareRepository model.FareRepositoryInterface
}

type FareResult struct {
	Service         string `json:"service"`
	Price           string `json:"price"`
	DaysForDelivery string `json:"days_for_delivery"`
	DeliverHome     string `json:"deliver_home"`
	DeliverSaturday string `json:"deliver_saturday"`
	Obs             string `json:"obs"`
}

func (f *FareUseCase) GetFare(fare model.Fare) (FareResult, error) {
	fareResult := FareResult{}

	err := fare.IsValid()
	if err != nil {
		return fareResult, err
	}

	resultByte, err := f.FareRepository.GetFare(fare)
	if err != nil {
		return fareResult, err
	}

	err = json.Unmarshal(resultByte, &fareResult)
	if err != nil {
		return fareResult, err
	}

	return fareResult, nil
}
