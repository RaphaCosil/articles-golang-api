package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/service"
)

type ArticleHandler struct {
	service service.ArticleService
}

func NewArticleHandler(s service.ArticleService) *ArticleHandler {
	return &ArticleHandler{s}
}

func (h *ArticleHandler) FindAll(c *gin.Context) {
	articles, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error": "Error getting all articles"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) FindById(c *gin.Context) {
	articleId := c.Param("id")
	n, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	article, err := h.service.FindById(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error": "Error getting article by id"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdArticle, err := h.service.Create(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating article"})
		return
	}
	c.JSON(http.StatusCreated, createdArticle)
}

func (h *ArticleHandler) Update(c *gin.Context) {
	articleId := c.Param("id")
	n, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err = h.service.Update(n, article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating article"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully"})
}

func (h *ArticleHandler) DeleteById(c *gin.Context) {
	articleId := c.Param("id")
	n, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	err = h.service.DeleteById(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting article"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func (h *ArticleHandler) FindByCustomerId(c *gin.Context) {
	customerId := c.Param("customerId")
	n, err := strconv.Atoi(customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	articles, err := h.service.FindByCustomerId(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting articles by customer id"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) FindByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title query parameter is required"})
		return
	}
	articles, err := h.service.FindByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting articles by title"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) FindByContent(c *gin.Context) {
	content := c.Query("content")
	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content query parameter is required"})
		return
	}
	articles, err := h.service.FindByContent(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting articles by content"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) FindByKeywords(c *gin.Context) {
	keywords := c.QueryArray("keywords")
	if len(keywords) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Keywords query parameter is required"})
		return
	}
	articles, err := h.service.FindByKeywords(keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting articles by keywords"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) FindByKeywordsFilter(c *gin.Context) {
	keywords := c.QueryArray("keywords")
	if len(keywords) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Keywords query parameter is required"})
		return
	}
	articles, err := h.service.FindByKeywordsFilter(keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting articles by keywords filter"})
		return
	}
	c.JSON(http.StatusOK, articles)
}
