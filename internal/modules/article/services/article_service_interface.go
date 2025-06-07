package services

import ArticleModel "gin-demo/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []ArticleModel.Article
	GetStoriesArticles() []ArticleModel.Article
}
