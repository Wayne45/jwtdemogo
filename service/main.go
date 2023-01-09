package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jwtdemogo/controller"
	"github.com/jwtdemogo/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /v1
func main() {
	//r := gin.Default()
	r := gin.New() // Do not print request logs on server.

	docs.SwaggerInfo.BasePath = "/v1"
	v1 := r.Group("/v1")
	{
		v1.POST("/jwt-tokens", controller.SignJwt)
		v1.GET("/jwt-tokens", controller.VerifyJwt)

		v1.GET("/self", controller.ShowSelf)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":30001")
}
