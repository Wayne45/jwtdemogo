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
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/v1"
	v1 := r.Group("/v1")
	{
		v1.POST("/jwt-tokens", controller.SignJwt)
		v1.GET("/jwt-tokens", controller.VerifyJwt)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":30001")
}
