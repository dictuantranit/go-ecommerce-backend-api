package initialize

import (
	"github.com/dictuantranit/go-ecommerce-backend-api/global"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/database"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/service"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	//..................
}
