package routes

import (
	articleController "gin-demo/internal/modules/article/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articleController := articleController.New()
	router.GET("/articles/:id", articleController.Show)
}
