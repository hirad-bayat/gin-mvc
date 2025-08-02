package services

import (
	"gin-demo/internal/modules/user/requests/auth"
	UserResponses "gin-demo/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponses.User, error)
}
