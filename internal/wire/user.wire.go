//go:build wireinject

package wire

import (
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/controller"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/repo"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHanlder() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
