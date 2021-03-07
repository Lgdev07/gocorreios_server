package repository

import (
	"encoding/json"
	"testing"

	"github.com/Lgdev07/gocorreios/fare"
	"github.com/Lgdev07/gocorreios_server/domain/model"
	"github.com/stretchr/testify/assert"
)

type MockGoCorreiosRepository struct{}

func (mockRepo *MockGoCorreiosRepository) Fare(fareInterf fare.Interface) ([]byte, error) {
	print("passei aqui")

	jsonResponse := `{
		"service": "SEDEEEEEX",
		"price": "22,50",
		"days_for_delivery": "3",
		"deliver_home": "N",
		"deliver_saturday": "S",
		"obs": "O CEP de destino está temporariamente sem entrega domiciliar. A entrega será efetuada na agência indicada no Aviso de Chegada que será entregue no endereço do destinatário."
	}`

	return []byte(jsonResponse), nil
}

func TestGetFare(t *testing.T) {
	fareMock := GoCorreiosRepository{
		Repo: &MockGoCorreiosRepository{},
	}

	fareModel := model.Fare{
		CepDestination: "09981380",
		CepOrigin:      "09981380",
		Height:         "15",
		Lenght:         "15",
		Service:        "SEDEX",
		Weight:         "1",
		Width:          "15",
	}

	result, err := fareMock.GetFare(fareModel)
	if err != nil {
		t.Errorf("Erro getting the GetFare: %v", err)
	}

	resultInterface := make(map[string]interface{})

	err = json.Unmarshal(result, &resultInterface)
	if err != nil {
		t.Errorf("Erro Unmarshal json: %v", err)
	}

	assert.Equal(t, "SEDEEEEEX", resultInterface["service"])
	assert.Equal(t, "22,50", resultInterface["price"])
	assert.Equal(t, "3", resultInterface["days_for_delivery"])
	assert.Equal(t, "N", resultInterface["deliver_home"])
	assert.Equal(t, "S", resultInterface["deliver_saturday"])
}
