package services

import (
	"errors"
	"quasarproject/entities"
	"quasarproject/responses"
)

// TopSecretServiceInterface is an interface
type TopSecretServiceInterface interface {
	ParseMessage(info entities.ParseMessage) (responses.MessageInfo, error)
}

type topSecretService struct {
	calculatorService CalculatorServiceInterface
}

// NewTopSecretService implements service
func NewTopSecretService(calculatorService CalculatorServiceInterface) TopSecretServiceInterface {
	return &topSecretService{
		calculatorService: calculatorService,
	}
}

func (s topSecretService) ParseMessage(message entities.ParseMessage) (responses.MessageInfo, error) {

	var response responses.MessageInfo
	var distances []float32

	distances = append(distances, message.Distances[firstSatelite], message.Distances[secondSatelite], message.Distances[thirdSatelite])

	cordX, cordY := s.calculatorService.GetLocation(distances...)

	if cordX*cordY == returnError {
		return responses.MessageInfo{}, errors.New("Error al calcular las coordenadas")
	}

	obtainedMessage := s.calculatorService.GetMessage(message.Messages...)

	if obtainedMessage == "" {
		return responses.MessageInfo{}, errors.New("El mensaje recibido está vacío")
	}

	response.Message = obtainedMessage
	response.Position.X = cordX
	response.Position.Y = cordY

	return response, nil
}
