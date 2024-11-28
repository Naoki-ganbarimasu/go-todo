package usecase

import (
	"go-test-api/model"
	"go-test-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	GetUserByEmail(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if error != nil {
		return model.UserResponse{}, error
	}
}
