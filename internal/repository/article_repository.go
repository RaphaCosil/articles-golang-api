package repository

import (
	"github.com/raphacosil/go-study/internal/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll() ([]model.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) FindAll() ([]model.Article, error) {
	var articles []model.Article
	result := r.db.Find(&articles)
	return articles, result.Error
}
