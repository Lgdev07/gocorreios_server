package usecase

import (
	"encoding/json"

	"github.com/Lgdev07/gocorreios_server/domain/model"
)

type TrackingUseCase struct {
	TrackingRepository model.TrackingRepositoryInterface
}

type TrackingResult []map[string]interface{}

func (f *TrackingUseCase) GetTracking(tracking model.Tracking) (TrackingResult, error) {
	trackingResult := TrackingResult{}

	err := tracking.IsValid()
	if err != nil {
		return trackingResult, err
	}

	resultByte, err := f.TrackingRepository.GetTracking(tracking)
	if err != nil {
		return trackingResult, err
	}

	err = json.Unmarshal(resultByte, &trackingResult)
	if err != nil {
		return trackingResult, err
	}

	return trackingResult, nil
}
