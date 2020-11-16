package services

import (
	"log"
	"quasarproject/entities"
	"quasarproject/responses"
	"testing"
)

// TestParseMessage tests ParseMessage mehotd
func TestParseMessage(t *testing.T) {

	var message entities.ParseMessage
	var response responses.MessageInfo
	distances := make(map[string]float32)

	distances["kenobi"] = 627
	distances["skywalker"] = 208
	distances["sato"] = 450

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
	response.Position.X = float32(49)
	response.Position.Y = float32(103)

	expected := response

	if result != expected {
		log.Fatal(result, " es distinto a: ", expected)
	}
}
