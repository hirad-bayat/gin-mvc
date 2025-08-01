package auth

type RegisterRequest struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required,min=3,max=100"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required,email,min=3,max=100"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,min=8"`
}
