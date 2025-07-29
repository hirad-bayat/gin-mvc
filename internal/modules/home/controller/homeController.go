package controller

import (
	ArticleService "gin-demo/internal/modules/article/services"
	"gin-demo/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
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

func (controller *Controller) Articles(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"feature": controller.articleService.GetFeaturedArticles(),
		"stories": controller.articleService.GetStoriesArticles(),
	})
}
