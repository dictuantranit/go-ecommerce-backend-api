package main

import (
	"net/http"

	_ "github.com/dictuantranit/go-ecommerce-backend-api/cmd/swag/docs"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/initialize"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count_total",
		Help: "Total number of ping requests.",
	},
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping",
	})
}

func ping(c *gin.Context) {
	pingCounter.Inc() // 1 2 3
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

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

	prometheus.MustRegister(pingCounter)

	r.GET("/ping/200", ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8002")
}
