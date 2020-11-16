package controllers

import (
	"log"
	"net/http"
	"quasarproject/responses"
	"quasarproject/services"

	"github.com/gin-gonic/gin"
)

// ControllerInterface is an interface
type ControllerInterface interface {
	CheckPingHandler(g *gin.Context)
	ParseMessageHandler(g *gin.Context)
	PingDBHandler(g *gin.Context)
	ParseMessageSplitHandler(g *gin.Context)
	SaveParseMessageSplitHandler(g *gin.Context)
}

// NewControllers implement services
func NewControllers(
	paramsService services.ParametersServiceInterface,
	tsService services.TopSecretServiceInterface,
	statService services.StatusServiceInterface,
	tssService services.TopSecretServiceSplitInterface,
) ControllerInterface {
	return &controllers{
		parametersService:     paramsService,
		topSecretService:      tsService,
		statusService:         statService,
		topSecretSplitService: tssService,
	}
}

type controllers struct {
	parametersService     services.ParametersServiceInterface
	topSecretService      services.TopSecretServiceInterface
	statusService         services.StatusServiceInterface
	topSecretSplitService services.TopSecretServiceSplitInterface
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
	log.Print(parameters)
	response, err := c.topSecretService.ParseMessage(parameters)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusNotFound, nil)
		panic(err)
	}

	g.JSON(http.StatusOK, response)
}

func (c controllers) ParseMessageSplitHandler(ctx *gin.Context) {

	info, errSplit := c.topSecretSplitService.ObtainSateliteInfo()

	if errSplit != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.ErrorSplit{Error: "No hay suficiente información"})
		panic("error")
	}

	parseMessage, err := c.topSecretService.ParseMessage(info)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.ErrorSplit{Error: "No hay suficiente información"})
		panic("error")
	}

	ctx.JSON(http.StatusOK, parseMessage)
}

func (c controllers) SaveParseMessageSplitHandler(ctx *gin.Context) {

	info, errParams := c.parametersService.ObtainSplitParameters(ctx)

	if errParams != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, nil)
		panic("error")
	}

	satelite := info

	satelite.Name = ctx.Param("satelite")

	err := c.topSecretSplitService.SaveSateliteInfo(satelite)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, nil)
		panic("error")
	}

	ctx.JSON(http.StatusOK, nil)
}
