package main

import (
	"quasarproject/controllers"
	"quasarproject/http"
	"quasarproject/services"
)

var (
	httpRouter         http.RouterInterface                    = http.NewGinRouter()
	calculatorService  services.CalculatorServiceInterface     = services.NewCalculatorService()
	topSecretService   services.TopSecretServiceInterface      = services.NewTopSecretService(calculatorService)
	parametersService  services.ParametersServiceInterface     = services.NewGinParamsService()
	projectControllers controllers.ProjectControllersInterface = controllers.NewControllers(parametersService, topSecretService)
)

func main() {
	httpRouter.GET("/ping", projectControllers.CheckPingHandler)
	httpRouter.POST("/topsecret/", projectControllers.ParseMessageHandler)
	httpRouter.SERVE("9200")
}
