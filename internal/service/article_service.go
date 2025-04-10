package service

import(
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/repository"
)

type ArticleService interface {
	FindAll() ([]model.Article, error)
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(r repository.ArticleRepository) ArticleService {
	return &articleService{r}
}

func (s *articleService) FindAll() ([]model.Article, error) {
	return s.repo.FindAll()
}