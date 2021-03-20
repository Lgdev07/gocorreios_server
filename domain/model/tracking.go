package model

import (
	"github.com/asaskevich/govalidator"
)

type TrackingRepositoryInterface interface {
	GetTracking(trackingModel Tracking) ([]byte, error)
}

type Tracking struct {
	Codes []string `json:"codes" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (tracking *Tracking) IsValid() error {
	_, err := govalidator.ValidateStruct(tracking)
	if err != nil {
		return err
	}
	return nil
}
