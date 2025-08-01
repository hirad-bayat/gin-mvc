package routes

import (
	userController "gin-demo/internal/modules/user/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	userController := userController.New()
	router.GET("/login", userController.Login)
	router.GET("/register", userController.Register)
	router.POST("/register", userController.HandleRegister)

}
