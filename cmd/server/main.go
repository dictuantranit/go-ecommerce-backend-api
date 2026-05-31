package main

import (
	"net/http"

	"github.com/dictuantranit/go-ecommerce-backend-api/internal/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.Run()
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping",
	})
}
