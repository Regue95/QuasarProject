package http

import "github.com/gin-gonic/gin"

type gRouter struct {
}

var (
	router = gin.Default()
)

// NewGinRouter implements gRouter
func NewGinRouter() RouterInterface {
	return &gRouter{}
}

func (gr gRouter) GET(url string, i interface{}) {
	router.GET(url, i.(func(g *gin.Context)))
}

func (gr gRouter) POST(url string, i interface{}) {
	router.POST(url, i.(func(g *gin.Context)))
}

func (gr gRouter) SERVE(port string) {
	router.Run(":" + port)
}
