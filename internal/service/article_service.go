package service

import(
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/repository"
)

type ArticleService interface {
	FindAll() ([]model.Article, error)
	FindById(articleId int) (model.Article, error)
	Create(article model.Article) (*model.Article, *error)
	Update(articleId int, article model.Article) error
	DeleteById(articleId int) error
	FindByCustomerId(customerId int) ([]model.Article, error)
	FindByTitle(title string) ([]model.Article, error)
	FindByContent(content string) ([]model.Article, error)
	FindByKeywords(keywords []string) ([]model.Article, error)
	FindByKeywordsFilter(keywords []string) ([]model.Article, error)
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

func (s *articleService) FindById(articleId int) (model.Article, error) {
	return s.repo.FindById(articleId)
}

func (s *articleService) Create(article model.Article) (*model.Article, *error){
	return s.repo.Create(article)
}

func (s *articleService) Update(articleId int, article model.Article) error {
	return s.repo.Update(articleId, article)
}

func (s *articleService) DeleteById(articleId int) error {
	return s.repo.DeleteById(articleId)
}

func (s *articleService) FindByCustomerId(customerId int) ([]model.Article, error) {
	return s.repo.FindByCustomerId(customerId)
}

func (s *articleService) FindByTitle(title string) ([]model.Article, error) {
	return s.repo.FindByTitle(title)
}

func (s *articleService) FindByContent(content string) ([]model.Article, error) {
	return s.repo.FindByContent(content)
}

func (s *articleService) FindByKeywords(keywords []string) ([]model.Article, error) {
	return s.repo.FindByKeywords(keywords)
}

func (s *articleService) FindByKeywordsFilter(keywords []string) ([]model.Article, error) {
	return s.repo.FindByKeywordsFilter(keywords)
}
