package repository

import (
	"encoding/json"
	"testing"

	"github.com/Lgdev07/gocorreios/fare"
	"github.com/Lgdev07/gocorreios_server/domain/model"
	"github.com/stretchr/testify/assert"
)

type MockGoCorreiosRepository struct {
	TrackingResponse func() ([]byte, error)
	FareResponse     func() ([]byte, error)
}

func (mockRepo *MockGoCorreiosRepository) Fare(fareInterf fare.Interface) ([]byte, error) {
	return mockRepo.FareResponse()
}

func (mockRepo *MockGoCorreiosRepository) Tracking(codesList []string) ([]byte, error) {
	return mockRepo.TrackingResponse()
}

func TestGetFareSuccess(t *testing.T) {
	jsonResponse := `{
		"service": "SEDEEEEEX",
		"price": "22,50",
		"days_for_delivery": "3",
		"deliver_home": "N",
		"deliver_saturday": "S",
		"obs": "O CEP de destino está temporariamente sem entrega domiciliar. A entrega será efetuada na agência indicada no Aviso de Chegada que será entregue no endereço do destinatário."
	}`

	fareResult := func() ([]byte, error) {
		return []byte(jsonResponse), nil
	}

	fareMock := GoCorreiosRepository{
		Repo: &MockGoCorreiosRepository{
			FareResponse: fareResult,
		},
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

func TestGetFareError(t *testing.T) {
	jsonResponse := `{
		"error": "CEP de origem invalido."
	}`

	fareResult := func() ([]byte, error) {
		return []byte(jsonResponse), nil
	}

	fareMock := GoCorreiosRepository{
		Repo: &MockGoCorreiosRepository{
			FareResponse: fareResult,
		},
	}

	fareModel := model.Fare{
		CepDestination: "09981380",
		CepOrigin:      "099841380",
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

	assert.Equal(t, "CEP de origem invalido.", resultInterface["error"])
}

func TestGetTrackingSuccess(t *testing.T) {
	jsonResponse := `[
		{
		  "category": "SEDEX",
		  "error": "",
		  "history": [
			{
			  "date": "05/03/2021 - 11:24",
			  "description": "Objeto saiu para entrega ao destinatário",
			  "destination": "",
			  "detail": "",
			  "origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
			  "status": 1,
			  "type": "OEC"
			}
		  ],
		  "last_date": "05/03/2021 - 14:04",
		  "last_description": "Objeto entregue ao destinatário",
		  "last_destination": "",
		  "last_detail": "",
		  "last_origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
		  "last_status": 1,
		  "last_type": "BDE",
		  "number": "ON732404576BR"
		}
	  ]`

	trackingResult := func() ([]byte, error) {
		return []byte(jsonResponse), nil
	}

	trackingMock := GoCorreiosRepository{
		Repo: &MockGoCorreiosRepository{
			TrackingResponse: trackingResult,
		},
	}

	trackingModel := model.Tracking{
		Codes: []string{"ON732404576BR"},
	}

	result, err := trackingMock.GetTracking(trackingModel)
	if err != nil {
		t.Errorf("Erro getting the GetTracking: %v", err)
	}

	var resultInterface []map[string]interface{}

	err = json.Unmarshal(result, &resultInterface)
	if err != nil {
		t.Errorf("Erro Unmarshal json: %v", err)
	}

	assert.Equal(t, 1, len(resultInterface))
	assert.Equal(t, "05/03/2021 - 14:04", resultInterface[0]["last_date"])
	assert.Equal(t, "Objeto entregue ao destinatário", resultInterface[0]["last_description"])
	assert.Equal(t, float64(1), resultInterface[0]["last_status"])
	assert.Equal(t, "BDE", resultInterface[0]["last_type"])
	assert.Equal(t, "ON732404576BR", resultInterface[0]["number"])
}

func TestGetTrackingError(t *testing.T) {
	jsonResponse := `[
		{
		  "error": "ERRO: Objeto não encontrado na base de dados dos Correios.",
		  "number": "ON7329044576B"
		},
		{
		  "error": "ERRO: Objeto não encontrado na base de dados dos Correios.",
		  "number": "R"
		}
	  ]`

	trackingResult := func() ([]byte, error) {
		return []byte(jsonResponse), nil
	}

	trackingMock := GoCorreiosRepository{
		Repo: &MockGoCorreiosRepository{
			TrackingResponse: trackingResult,
		},
	}

	trackingModel := model.Tracking{
		Codes: []string{"ON7329044576BR"},
	}

	result, err := trackingMock.GetTracking(trackingModel)
	if err != nil {
		t.Errorf("Erro getting the GetTracking: %v", err)
	}

	var resultInterface []map[string]interface{}

	err = json.Unmarshal(result, &resultInterface)
	if err != nil {
		t.Errorf("Erro Unmarshal json: %v", err)
	}

	assert.Equal(t, 2, len(resultInterface))
	assert.Equal(t, "ERRO: Objeto não encontrado na base de dados dos Correios.", resultInterface[0]["error"])
	assert.Equal(t, "ON7329044576B", resultInterface[0]["number"])
	assert.Equal(t, "ERRO: Objeto não encontrado na base de dados dos Correios.", resultInterface[1]["error"])
	assert.Equal(t, "R", resultInterface[1]["number"])
}
