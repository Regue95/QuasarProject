package services

import (
	"log"
	"quasarproject/entities"
	"quasarproject/responses"
	"testing"
)

type calculationServiceMock struct {
}

// TestParseMessage tests ParseMessage mehotd
func TestParseMessage(t *testing.T) {

	var message entities.ParseMessage
	var response responses.MessageInfo
	distances := make(map[string]float32)

	distances["kenobi"] = 722
	distances["skywalker"] = 300
	distances["sato"] = 410

	message.Messages = [][]string{
		{" ", "es", "", " ", ""},
		{"este", "", "un", "", "secreto"},
		{"", "", "", "mensaje", ""},
	}

	message.Distances = distances

	calculationService := NewCalculatorService()

	service := NewTopSecretService(calculationService)

	result, _ := service.ParseMessage(message)

	response.Message = "este es un mensaje secreto"
	response.Position.X = float32(100)
	response.Position.Y = float32(204)

	expected := response

	if result != expected {
		log.Fatal(result, " es distinto a: ", expected)
	}
}
