package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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