package service

import(
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/repository"
)

type KeywordService interface {
	FindAll() ([]model.Key_word, error)
	FindById(keywordId int) (model.Key_word, error)
	Create(keyword model.Key_word) (*model.Key_word, *error)
	Update(keywordId int, keyword model.Key_word) error
	DeleteById(keywordId int) error
	FindByArticleId(articleId int) ([]model.Article, error)
}

type keywordService struct {
	repo repository.KeywordRepository
}

func NewKeywordService(r repository.KeywordRepository) KeywordService {
	return &keywordService{r}
}

func (s *keywordService) FindAll() ([]model.Key_word, error) {
	return s.repo.FindAll()
}

func (s *keywordService) FindById(keywordId int) (model.Key_word, error) {
	return s.repo.FindById(keywordId)
}

func (s *keywordService) Create(keyword model.Key_word) (*model.Key_word, *error) {
	return s.repo.Create(keyword)
}

func (s *keywordService) Update(keywordId int, keyword model.Key_word) error {
	return s.repo.Update(keywordId, keyword)
}

func (s *keywordService) DeleteById(keywordId int) error {
	return s.repo.DeleteById(keywordId)
}

func (s *keywordService) FindByArticleId(articleId int) ([]model.Article, error) {
	return s.repo.FindByArticleId(articleId)
}