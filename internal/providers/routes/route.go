package routes

import (
	articleRoutes "gin-demo/internal/modules/article/routes"
	homeRoutes "gin-demo/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)

}
