package repositories

import ArticleModel "gin-demo/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
}
