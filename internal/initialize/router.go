package initialize

import (
	c "github.com/dictuantranit/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}

	return r
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware

	//...
	return r
}
