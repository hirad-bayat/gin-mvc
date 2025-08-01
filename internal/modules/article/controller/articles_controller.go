package controller

import (
	ArticleService "gin-demo/internal/modules/article/services"
	"gin-demo/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	//get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the id"})
		//return
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/error", gin.H{"title": "Server Error", "message": "error converting the id", "status": http.StatusInternalServerError})
		return
	}

	//find the article from database
	article, err := controller.articleService.Find(id)

	//if the article is not found show an error page
	if err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		//return
		html.Render(c, http.StatusNotFound, "templates/errors/html/error", gin.H{"title": "Page Not Found", "message": err.Error(), "status": http.StatusNotFound})
		return
	}

	//if article found, render article template
	c.JSON(http.StatusOK, gin.H{"article": article})
}
