package services

import (
	"encoding/json"
	"log"
	"quasarproject/entities"
	"quasarproject/repositories"
)

// TopSecretServiceSplitInterface is an interface
type TopSecretServiceSplitInterface interface {
	ObtainSateliteInfo() (entities.ParseMessage, error)
	SaveSatelite(data entities.Satelites) error
}

type topSecretServiceSplit struct {
	calculatorService CalculatorServiceInterface
	repository        repositories.RepositoryInterface
}

// NewTopSecretServiceSplit implements TopSecretServiceSplitInterface
func NewTopSecretServiceSplit(calculatorService CalculatorServiceInterface, repository repositories.RepositoryInterface) TopSecretServiceSplitInterface {
	return &topSecretServiceSplit{
		calculatorService: calculatorService,
		repository:        repository,
	}
}

func (ss topSecretServiceSplit) ObtainSateliteInfo() (entities.ParseMessage, error) {
	var parseMessage entities.ParseMessage
	var satelitesInfo entities.Satelite
	var messages [][]string
	distances := make(map[string]float32)

	sateliteInfo, err := ss.repository.Get("satelite")

	if err != nil {
		return parseMessage, err
	}

	_ = json.Unmarshal([]byte(sateliteInfo), &satelitesInfo)

	for _, info := range satelitesInfo.Satelites {
		distances[info.Name] = info.Distance
		messages = append(messages, info.Message)
	}

	parseMessage.Distances = distances
	parseMessage.Messages = messages

	return parseMessage, nil
}

func (ss topSecretServiceSplit) SaveSatelite(info entities.Satelites) error {
	var satelitesInfo entities.Satelite

	setSatelite, _ := json.Marshal(satelitesInfo)

	err := ss.repository.Set("satelite", string(setSatelite))

	if err != nil {
		log.Print("Error al guardar satelite. Error: ", err.Error())
		return err
	}

	return nil
}
