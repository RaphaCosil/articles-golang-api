package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

