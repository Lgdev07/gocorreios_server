package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Fare struct {
	CepDestination string `json:"cep_destination"`
	CepOrigin      string `json:"cep_origin"`
	Height         string `json:"height"`
	Lenght         string `json:"lenght"`
	Service        string `json:"service"`
	Weight         string `json:"weight"`
	Width          string `json:"width"`
}

func (f Fare) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.CepDestination, validation.Required),
		validation.Field(&f.CepOrigin, validation.Required),
		validation.Field(&f.Height, validation.Required),
		validation.Field(&f.Lenght, validation.Required),
		validation.Field(&f.Service, validation.Required),
		validation.Field(&f.Weight, validation.Required),
		validation.Field(&f.Width, validation.Required),
	)
}
