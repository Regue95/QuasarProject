package services

import (
	"encoding/json"
	"quasarproject/entities"
	"quasarproject/repositories"
)

// TopSecretServiceSplitInterface is an interface
type TopSecretServiceSplitInterface interface {
	ObtainSateliteInfo() (entities.ParseMessage, error)
	SaveSateliteInfo(data entities.Satelites) error
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
	var messages [][]string
	aux := make(map[string]entities.SplitMessageData)
	distances := make(map[string]float32)

	sateliteInfo, err := ss.repository.Get("satelite")

	if err != nil {
		return parseMessage, err
	}

	_ = json.Unmarshal([]byte(sateliteInfo), &aux)

	for k, v := range aux {
		distances[k] = v.Distance
		messages = append(messages, v.Message)

	}
	parseMessage.Distances = distances
	parseMessage.Messages = messages

	return parseMessage, nil
}

func (ss topSecretServiceSplit) SaveSateliteInfo(info entities.Satelites) error {
	aux := make(map[string]entities.SplitMessageData)

	getResponse, _ := ss.repository.Get("satelite")

	if getResponse == "" {
		ss.saveSatelite(info)
	}

	_ = json.Unmarshal([]byte(getResponse), &aux)

	_, ok := aux[info.Name]

	if ok {
		return nil
	}

	data := entities.SplitMessageData{
		Message:  info.Message,
		Distance: info.Distance,
	}

	aux[info.Name] = data

	add, _ := json.Marshal(aux)

	setErr := ss.repository.Set("satelite", string(add))

	if setErr != nil {
		return setErr
	}

	return nil
}

func (ss topSecretServiceSplit) saveSatelite(info entities.Satelites) error {
	sateliteInf := make(map[string]entities.SplitMessageData)

	data := entities.SplitMessageData{
		Message:  info.Message,
		Distance: info.Distance,
	}

	sateliteInf[info.Name] = data

	addSat, _ := json.Marshal(sateliteInf)

	setErr := ss.repository.Set("satelite", string(addSat))

	if setErr != nil {
		return setErr
	}

	return nil
}
