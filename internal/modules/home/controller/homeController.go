package controller

import (
	"gin-demo/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Home Page",
	})
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/login", gin.H{
		"title": "Login Page",
	})
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/register", gin.H{
		"title": "Register Page",
	})
}
