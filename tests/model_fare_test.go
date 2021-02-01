package tests

import (
	"testing"

	"github.com/Lgdev07/gocorreios_server/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateFare(t *testing.T) {
	fare := models.Fare{
		CepDestination: "09981380",
		CepOrigin:      "09981380",
		Height:         "15",
		Lenght:         "15",
		Service:        "10",
		Weight:         "11",
		Width:          "",
	}

	assert.Error(t, fare.Validate())

	fare.Width = "10"

	assert.Nil(t, fare.Validate())
}
