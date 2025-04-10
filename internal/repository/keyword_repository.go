package repository

import (
	"github.com/raphacosil/go-study/internal/model"
	"gorm.io/gorm"
)

type KeywordRepository interface {
	FindAll() ([]model.Key_word, error)
	FindById(keywordId int) (model.Key_word, error)
	Create(keyword model.Key_word) (*model.Key_word, *error)
	Update(keywordId int, keyword model.Key_word) error
	DeleteById(keywordId int) error
	FindByArticleId(articleId int) ([]model.Article, error)
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

func (r *keywordRepository) FindById(keywordId int) (model.Key_word, error) {
	var keyword model.Key_word
	result := r.db.First(&keyword, keywordId)
	return keyword, result.Error
}

func (r *keywordRepository) Create(keyword model.Key_word) (*model.Key_word, *error) {
	err := r.db.Create(&keyword).Error
	if err !=nil {
		return nil, &err
	}
	return &keyword, nil
}

func (r *keywordRepository) Update(keywordId int, keyword model.Key_word) error {
	return r.db.Model(&model.Key_word{}).Where("key_word_id = ?", keywordId).Updates(keyword).Error
}

func (r *keywordRepository) DeleteById(keywordId int) error {
	return r.db.Delete(&model.Article{}, keywordId).Error
}

func (r *keywordRepository) FindByArticleId(articleId int) ([]model.Article, error) {
	var keywords []model.Article
	result := r.db.Where("article_id = ?", articleId).Find(&keywords)
	return keywords, result.Error
}
