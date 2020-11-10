package services

import (
	"errors"
	"quasarproject/entities"

	"github.com/gin-gonic/gin"
)

const (
	firstSatelite  = "kenobi"
	secondSatelite = "skywalker"
	thirdSatelite  = "sato"
	totalSatelites = 3
)

// ParametersServiceInterface is an interface
type ParametersServiceInterface interface {
	ObtainParameters(parameters interface{}) (entities.ParseMessage, error)
}

type gParamsService struct {
}

// NewGinParamsService implements ParametersServiceInterface
func NewGinParamsService() ParametersServiceInterface {
	return &gParamsService{}
}

func (g *gParamsService) ObtainParameters(info interface{}) (entities.ParseMessage, error) {
	var parseMessage entities.ParseMessage
	var sateliteInfo entities.Satelite
	var messages [][]string
	distances := make(map[string]float32)
	ctx := info.(*gin.Context)

	if err := ctx.ShouldBindJSON(&sateliteInfo); err != nil {
		return entities.ParseMessage{}, err
	}

	for _, info := range sateliteInfo.Satelites {
		distances[info.Name] = info.Distance
		messages = append(messages, info.Message)
	}

	parseMessage.Distances = distances
	parseMessage.Messages = messages

	if len(parseMessage.Messages) != totalSatelites {
		return entities.ParseMessage{}, errors.New("La cantidad de satélites ingresados es incorrecta")
	}

	return parseMessage, nil
}
