package routes

import (
	homeRoutes "gin-demo/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)

}
