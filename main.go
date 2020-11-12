package main

import (
	"os"
	"quasarproject/controllers"
	"quasarproject/database"
	"quasarproject/http"
	"quasarproject/repositories"
	"quasarproject/services"

	"github.com/joho/godotenv"
)

var (
	httpRouter         http.RouterInterface                = http.NewGinRouter()
	dbclient           database.ClientInterface            = database.NewClient()
	repository         repositories.RepositoryInterface    = repositories.NewRepository(dbclient)
	calculatorService  services.CalculatorServiceInterface = services.NewCalculatorService()
	topSecretService   services.TopSecretServiceInterface  = services.NewTopSecretService(calculatorService)
	parametersService  services.ParametersServiceInterface = services.NewGinParamsService()
	statusService      services.StatusServiceInterface     = services.NewStatusService(repository)
	projectControllers controllers.ControllerInterface     = controllers.NewControllers(parametersService, topSecretService, statusService)
)

func main() {
	_ = godotenv.Load()
	httpRouter.GET("/ping", projectControllers.CheckPingHandler)
	httpRouter.POST("/topsecret/", projectControllers.ParseMessageHandler)
	httpRouter.GET("/db-ping", projectControllers.PingDBHandler)
	httpRouter.SERVE(os.Getenv("PORT"))

}
