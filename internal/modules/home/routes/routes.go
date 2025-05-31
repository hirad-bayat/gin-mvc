package route

import (
	"gin-demo/pkg/html"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})

	router.GET("/sample", func(c *gin.Context) {
		c.HTML(http.StatusOK, "modules/home/html/home", gin.H{
			"title":    "Main website",
			"APP_NAME": viper.Get("App.Name"),
		})
	})

	router.GET("/", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home Page",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About Page",
		})
	})
}
