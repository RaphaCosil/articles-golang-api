package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error": "Error getting all users"})
		return
	}
	c.JSON(http.StatusOK, users)
}