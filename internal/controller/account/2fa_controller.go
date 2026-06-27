package account

import (
	"log"

	"github.com/dictuantranit/go-ecommerce-backend-api/internal/model"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/service"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/utils/context"
	"github.com/dictuantranit/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct{}

// User Setup Two Factor Authentication
// @Summary      Setup Two Factor Authentication
// @Description  Setup Two Factor Authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Authorization token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/setup [post]
func (sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid setupTwoFactorAuth parameter")
		return
	}
	// get UserId from uuid (token)
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		//response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "UserId is not invalid")
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
	}

	log.Println("UserId: ", userId)
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}

// User Verify Two Factor Authentication
// @Summary      Verify Two Factor Authentication
// @Description  Verify Two Factor Authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Authorization token"
// @Param        payload body model.TwoFactorVerificationInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/verify [post]
func (sUser2FA) VerifyTwoFactorAuth(ctx *gin.Context) {
	var params model.TwoFactorVerificationInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, "Missing or invalid params")
		return
	}

	// get UserId from uuid (token)
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, err.Error())
		//response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, "UserId is not invalid")
		return
	}
	log.Println("UserId: ", userId)
	params.UserId = uint32(userId)

	codeResult, err := service.UserLogin().VerifyTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}
