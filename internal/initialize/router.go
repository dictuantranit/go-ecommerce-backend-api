package initialize

import (
	"github.com/dictuantranit/go-ecommerce-backend-api/global"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/middlewares"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global
	r.Use(middlewares.ValidatorMiddleware()) // middleware

	// r.Use(middlewares.NewRateLimiter().GlobalRateLimiter()) // 100 req/s
	r.GET("/ping/100", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 100",
		})
	})

	// r.Use(middlewares.NewRateLimiter().PublicAPIRateLimiter()) // 80 req/s
	r.GET("/ping/80", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 80",
		})
	})

	// r.Use(middlewares.NewRateLimiter().UserAndPrivateRateLimiter()) // 50 req/s
	r.GET("/ping/50", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong 50",
		})
	})

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2026")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}
	return r
}
