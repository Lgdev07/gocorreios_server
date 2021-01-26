package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Lgdev07/gocorreios"
	"github.com/Lgdev07/gocorreios/fare"
	"github.com/Lgdev07/gocorreios_server/models"
	"github.com/gofiber/fiber/v2"
)

type FareResult struct {
	Service         string `json:"service"`
	Price           string `json:"price"`
	DaysForDelivery string `json:"days_for_delivery"`
	DeliverHome     string `json:"deliver_home"`
	DeliverSaturday string `json:"deliver_saturday"`
	Obs             string `json:"obs"`
}

func ShowFare(c *fiber.Ctx) error {
	fareModel := &models.Fare{}
	if err := c.BodyParser(fareModel); err != nil {
		return err
	}
	if err := fareModel.Validate(); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	newWeight, _ := strconv.ParseFloat(fareModel.Weight, 64)
	newLength, _ := strconv.ParseFloat(fareModel.Lenght, 64)
	newHeight, _ := strconv.ParseFloat(fareModel.Height, 64)
	newWidth, _ := strconv.ParseFloat(fareModel.Width, 64)

	params := fare.Interface{
		Service:        fareModel.Service,
		CepOrigin:      fareModel.CepOrigin,
		CepDestination: fareModel.CepDestination,
		Weight:         newWeight,
		Length:         newLength,
		Height:         newHeight,
		Width:          newWidth,
	}

	result, err := gocorreios.Fare(params)
	if err != nil {
		return err
	}

	var fareResult FareResult
	if err := json.Unmarshal(result, &fareResult); err != nil {
		return err
	}

	fmt.Println(string(result))
	return c.Status(fiber.StatusOK).JSON(fareResult)
}
