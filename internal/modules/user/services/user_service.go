package services

import (
	"errors"
	userModel "gin-demo/internal/modules/user/models"
	UserRepository "gin-demo/internal/modules/user/repositories"
	"gin-demo/internal/modules/user/requests/auth"
	UserResponses "gin-demo/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponses.User, error) {
	var response UserResponses.User
	var user userModel.User

	user.Name = request.Name
	user.Email = request.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		log.Fatal("hash password error")
		return response, errors.New("hash password error")
	}
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating user")
	}
	return UserResponses.ToUser(newUser), nil
}
