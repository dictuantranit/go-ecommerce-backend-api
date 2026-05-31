package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dictuantranit/go-ecommerce-backend-api/internal/repo"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/utils/random"
	"github.com/dictuantranit/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/dictuantranit/go-ecommerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
	//...
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
	//...
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s", hashEmail)

	// 1. check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. new OTP -> ...
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("Otp is :::%d\n", otp)

	// 3. save OTP in Redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	// fmt.Printf("err is :::%d\n", err)
	if err != nil {
		return response.ErrInvalidOTP
	}

	// 4. send Email OTP
	err = sendto.SendTemplateEmailOtp([]string{email}, "tttuan9th@gmail.com", "otp-auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	fmt.Printf("err sendto :::%d\n", err)
	if err != nil {
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}
