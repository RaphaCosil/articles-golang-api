package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/service"
)

type KeywordHandler struct {
	service service.KeywordService
}

func NewKeywordHandler(s service.KeywordService) *KeywordHandler {
	return &KeywordHandler{s}
}

func (h *KeywordHandler) FindAll(c *gin.Context) {
	keywords, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error": "Error getting all keywords"})
		return
	}
	c.JSON(http.StatusOK, keywords)
}

func (h *KeywordHandler) FindById(c *gin.Context) {
	keywordId := c.Param("id")
	n, err := strconv.Atoi(keywordId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid keyword ID"})
		return
	}
	keyword, err := h.service.FindById(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error": "Error getting keyword by id"})
		return
	}
	c.JSON(http.StatusOK, keyword)
}

func (h *KeywordHandler) Create(c *gin.Context) {
	var keyword model.Key_word
	if err := c.ShouldBindJSON(&keyword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdKeyword, err := h.service.Create(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating keyword"})
		return
	}
	c.JSON(http.StatusCreated, createdKeyword)
}

func (h *KeywordHandler) Update(c *gin.Context) {
	keywordId := c.Param("id")
	n, err := strconv.Atoi(keywordId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid keyword ID"})
		return
	}
	var keyword model.Key_word
	if err := c.ShouldBindJSON(&keyword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err = h.service.Update(n, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating keyword"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Keyword updated successfully"})
}

func (h *KeywordHandler) DeleteById(c *gin.Context) {
	keywordId := c.Param("id")
	n, err := strconv.Atoi(keywordId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid keyword ID"})
		return
	}
	err = h.service.DeleteById(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting keyword"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Keyword deleted successfully"})
}

func (h *KeywordHandler) FindByArticleId(c *gin.Context) {
	articleId := c.Param("articleId")
	n, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	articles, err := h.service.FindByArticleId(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting articles by customer id"})
		return
	}
	c.JSON(http.StatusOK, articles)
}
