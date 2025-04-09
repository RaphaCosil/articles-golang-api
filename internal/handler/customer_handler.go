package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(s service.CustomerService) *CustomerHandler {
	return &CustomerHandler{s}
}

func (h *CustomerHandler) FindAll(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error": "Error getting all users"})
		return
	}
	c.JSON(http.StatusOK, users)
}