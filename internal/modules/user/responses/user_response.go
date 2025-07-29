package responses

import (
	"fmt"
	UserModel "gin-demo/internal/modules/user/models"
)

type User struct {
	ID    uint
	Name  string
	Email string
	Role  string
	Image string
}

type Users struct {
	Data []User
}

func ToUser(user UserModel.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
