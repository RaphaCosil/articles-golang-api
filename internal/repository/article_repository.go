package repository

import (
	"github.com/raphacosil/go-study/internal/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
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

func (r *articleRepository) FindById(articleId int) (model.Article, error) {
	var article model.Article
	result := r.db.First(&article, articleId)
	return article, result.Error
}

func (r *articleRepository) Create(article model.Article) (*model.Article, *error){
	err := r.db.Create(&article).Error
	if err !=nil {
		return nil, &err
	}
	return &article, nil
}

func (r *articleRepository) Update(articleId int, article model.Article) error {
	return r.db.Model(&model.Article{}).Where("article_id = ?", articleId).Updates(article).Error
}

func (r *articleRepository) DeleteById(articleId int) error {
	return r.db.Delete(&model.Article{}, articleId).Error
}

func (r *articleRepository) FindByCustomerId(customerId int) ([]model.Article, error) {
	var articles []model.Article
	result := r.db.Where("sender_id = ?", customerId).Find(&articles)
	return articles, result.Error
}

func (r *articleRepository) FindByTitle(title string) ([]model.Article, error) {
	var articles []model.Article
	result := r.db.Where("title ILIKE ?", "%" + title + "%").Find(&articles)
	return articles, result.Error
}

func (r *articleRepository) FindByContent(content string) ([]model.Article, error) {
	var articles []model.Article
	result := r.db.Where("content ILIKE ?", "%" + content + "%").Find(&articles)
	return articles, result.Error
}

func (r *articleRepository) FindByKeywords(keywords []string) ([]model.Article, error) {
	var articles []model.Article

	sub := r.db.
	Table("article").
	Select("article.article_id").
	Joins("JOIN key_word ON key_word.article_id = article.article_id").
	Where("key_word.content IN ?", keywords)

	err := r.db.
	Where("article_id IN (?)", sub).
	Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}


func (r *articleRepository) FindByKeywordsFilter(keywords []string) ([]model.Article, error) {
	var articleIDs []uint

	err := r.db.
		Table("key_word").
		Select("article_id").
		Where("content IN ?", keywords).
		Group("article_id").
		Having("COUNT(DISTINCT content) = ?", len(keywords)).
		Pluck("article_id", &articleIDs).Error

	if err != nil {
		return nil, err
	}

	var articles []model.Article
	if len(articleIDs) == 0 {
		return []model.Article{}, nil
	}

	err = r.db.Where("article_id IN ?", articleIDs).Find(&articles).Error
	return articles, err
}
