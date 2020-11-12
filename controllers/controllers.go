package controllers

import (
	"net/http"
	"quasarproject/services"

	"github.com/gin-gonic/gin"
)

// ControllerInterface is an interface
type ControllerInterface interface {
	CheckPingHandler(g *gin.Context)
	ParseMessageHandler(g *gin.Context)
	PingDBHandler(g *gin.Context)
}

// NewControllers implement services
func NewControllers(
	paramsService services.ParametersServiceInterface,
	tsService services.TopSecretServiceInterface,
	statService services.StatusServiceInterface,
) ControllerInterface {
	return &controllers{
		parametersService: paramsService,
		topSecretService:  tsService,
		statusService:     statService,
	}
}

type controllers struct {
	parametersService services.ParametersServiceInterface
	topSecretService  services.TopSecretServiceInterface
	statusService     services.StatusServiceInterface
}

func (c controllers) CheckPingHandler(g *gin.Context) {
	response, err := c.statusService.GetPing()
	if err != nil {
		g.AbortWithStatusJSON(http.StatusNotFound, nil)
		panic(err)
	}
	g.JSON(http.StatusOK, response)
}

func (c controllers) PingDBHandler(g *gin.Context) {
	g.JSON(http.StatusOK, c.statusService.PingDB())
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
