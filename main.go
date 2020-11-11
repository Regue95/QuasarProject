package main

import (
	"os"
	"quasarproject/controllers"
	"quasarproject/http"
	"quasarproject/services"

	"github.com/joho/godotenv"
)

var (
	httpRouter         http.RouterInterface                    = http.NewGinRouter()
	calculatorService  services.CalculatorServiceInterface     = services.NewCalculatorService()
	topSecretService   services.TopSecretServiceInterface      = services.NewTopSecretService(calculatorService)
	parametersService  services.ParametersServiceInterface     = services.NewGinParamsService()
	projectControllers controllers.ProjectControllersInterface = controllers.NewControllers(parametersService, topSecretService)
)

func main() {
	_ = godotenv.Load()
	httpRouter.GET("/ping", projectControllers.CheckPingHandler)
	httpRouter.POST("/topsecret/", projectControllers.ParseMessageHandler)
	httpRouter.SERVE(os.Getenv("DB_PORT"))
}
