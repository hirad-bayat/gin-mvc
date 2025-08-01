package controller

import (
	"gin-demo/internal/modules/user/requests/auth"
	"gin-demo/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login Page",
	})
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register Page",
	})
}
func (controller *Controller) HandleRegister(c *gin.Context) {
	//validate the request
	var request auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Redirect(http.StatusFound, "/register")
		return
	}

	//create the user

	//check if there is any error on the user creation

	//after creating user redirect to home page

	c.JSON(http.StatusOK, gin.H{"message": "register handler"})
}
