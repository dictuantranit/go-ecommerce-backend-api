package service

import "github.com/dictuantranit/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user service u
func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
