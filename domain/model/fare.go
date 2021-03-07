package model

import (
	"github.com/asaskevich/govalidator"
)

type FareRepositoryInterface interface {
	GetFare(fareModel Fare) ([]byte, error)
}

type Fare struct {
	CepDestination string `json:"cep_destination" valid:"notnull"`
	CepOrigin      string `json:"cep_origin" valid:"notnull"`
	Height         string `json:"height" valid:"notnull"`
	Lenght         string `json:"lenght" valid:"notnull"`
	Service        string `json:"service" valid:"notnull"`
	Weight         string `json:"weight" valid:"notnull"`
	Width          string `json:"width" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (fare *Fare) IsValid() error {
	_, err := govalidator.ValidateStruct(fare)
	if err != nil {
		return err
	}
	return nil
}
