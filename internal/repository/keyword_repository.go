package repository

import (
	"github.com/raphacosil/go-study/internal/model"
	"gorm.io/gorm"
)

type KeywordRepository interface {
	FindAll() ([]model.Key_word, error)
}

type keywordRepository struct {
	db *gorm.DB
}

func NewKeywordRepository(db *gorm.DB) KeywordRepository {
	return &keywordRepository{db}
}

func (r *keywordRepository) FindAll() ([]model.Key_word, error) {
	var keywords []model.Key_word
	result := r.db.Find(&keywords)
	return keywords, result.Error
}
