package services

import (
	"errors"
	ArticleRepository "gin-demo/internal/modules/article/repositories"
	ArticleResponses "gin-demo/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (ArticleService *ArticleService) GetFeaturedArticles() ArticleResponses.Articles {
	articles := ArticleService.articleRepository.List(4)
	return ArticleResponses.ToArticles(articles)
}
func (ArticleService *ArticleService) GetStoriesArticles() ArticleResponses.Articles {
	articles := ArticleService.articleRepository.List(6)
	return ArticleResponses.ToArticles(articles)
}

func (ArticleService *ArticleService) Find(id int) (ArticleResponses.Article, error) {
	var response ArticleResponses.Article

	article := ArticleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	return ArticleResponses.ToArticle(article), nil
}
