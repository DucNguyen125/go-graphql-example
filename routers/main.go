package routers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/graphql", GraphqlHandler())
	if os.Getenv("GO_ENV") != "production" {
		router.GET("/graphql", PlaygroundHandler())
	}

	return router
}
