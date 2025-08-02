package repositories

import (
	userModel "gin-demo/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
}
