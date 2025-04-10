package service

import(
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/repository"
)

type KeywordService interface {
	FindAll() ([]model.Key_word, error)
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