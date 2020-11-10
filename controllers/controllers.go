package controllers

import (
	"net/http"
	"quasarproject/services"

	"github.com/gin-gonic/gin"
)

// ProjectControllersInterface is an interface
type ProjectControllersInterface interface {
	ControllerInterface
}

// ControllerInterface is an interface
type ControllerInterface interface {
	CheckPingHandler(g *gin.Context)
	ParseMessageHandler(g *gin.Context)
}

// NewControllers implement services
func NewControllers(
	paramsService services.ParametersServiceInterface,
	tsService services.TopSecretServiceInterface,
) ProjectControllersInterface {
	return &controllers{
		parametersService: paramsService,
		topSecretService:  tsService,
	}
}

type controllers struct {
	parametersService services.ParametersServiceInterface
	topSecretService  services.TopSecretServiceInterface
}

func (c controllers) CheckPingHandler(g *gin.Context) {
	g.JSON(http.StatusOK, nil)
}

func (c controllers) ParseMessageHandler(g *gin.Context) {

	parameters, err := c.parametersService.ObtainParameters(g)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusNotFound, nil)
		panic(err)
	}

	response, err := c.topSecretService.ParseMessage(parameters)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusNotFound, nil)
		panic(err)
	}

	g.JSON(http.StatusOK, response)
}
