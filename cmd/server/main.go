package main

import (
	"net/http"

	_ "github.com/dictuantranit/go-ecommerce-backend-api/cmd/swag/docs"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/initialize"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Ecommerce Backend SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/dictuantranit/go-ecommerce-backend-go

// @contact.name   TEAM GO
// @contact.url    github.com/dictuantranit/go-ecommerce-backend-go
// @contact.email  dictuantran.it@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2026
// @schema http
func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping",
	})
}
