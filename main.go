package main

import (
	"swagger_used/api"
	"swagger_used/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger Example Api"
	docs.SwaggerInfo.Description = "This is a sample server"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.Default()
	r.GET("/test", api.Hello) //golang规定

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8060")
}
