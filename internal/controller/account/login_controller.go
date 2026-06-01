package account

import (
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/service"
	"github.com/dictuantranit/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// management controller Login User
var Login = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	// Implement logic for login
	err := service.UserLogin().Login(ctx)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)

}

func (c *cUserLogin) Register(ctx *gin.Context) {
	var params 
}
