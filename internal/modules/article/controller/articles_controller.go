package controller

import (
	ArticleService "gin-demo/internal/modules/article/services"
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error converting the id"})
	}

	//find the article from database
	article, err := controller.articleService.Find(id)

	//if the article is not found show an error page
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	//if article found, render article template
	c.JSON(http.StatusOK, gin.H{"article": article})
}
