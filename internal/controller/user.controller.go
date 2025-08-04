package controller

import (
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/service"
	"github.com/dictuantranit/go-ecommerce-backend-api/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {

	response.SuccessResponse(c, 20001, []string{"Js", "cr7"})
	//response.ErrorResponse(c, 20003, "No need!")
}
