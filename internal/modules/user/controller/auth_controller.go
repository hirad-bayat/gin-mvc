package controller

import (
	"gin-demo/internal/modules/user/requests/auth"
	UserService "gin-demo/internal/modules/user/services"
	"gin-demo/pkg/errors"
	"gin-demo/pkg/html"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
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

		errors.Init()
		errors.SetFormErrors(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Get()})
		return

		c.Redirect(http.StatusFound, "/register")
		return
	}

	//create the user
	user, err := controller.userService.Create(request)

	//check if there is any error on the user creation
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	//after creating user redirect to home page
	log.Printf("user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}
