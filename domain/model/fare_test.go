package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel_FareIsValid(t *testing.T) {
	fare := Fare{
		CepDestination: "09998901",
		CepOrigin:      "09998901",
		Height:         "23",
		Lenght:         "53",
		Service:        "SEDEX",
		Weight:         "56",
	}

	err := fare.IsValid()
	assert.Error(t, err)

	fare.Width = "34"

	err = fare.IsValid()
	assert.Nil(t, err)
}
