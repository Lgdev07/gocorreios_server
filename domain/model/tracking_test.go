package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel_TrackingIsValid(t *testing.T) {
	tracking := Tracking{}

	err := tracking.IsValid()
	assert.Error(t, err)

	tracking.Codes = []string{"Test"}

	err = tracking.IsValid()
	assert.Nil(t, err)
}
