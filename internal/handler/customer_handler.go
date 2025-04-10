package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raphacosil/go-study/internal/model"
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

func (h *CustomerHandler) FindById(c *gin.Context) {

	customerId := c.Param("id")
	n, err := strconv.Atoi(customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	customer, err := h.service.FindById(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {"error": "Error getting customer by id"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdCustomer, err := h.service.Create(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating customer"})
		return
	}
	c.JSON(http.StatusCreated, createdCustomer)
}

func (h *CustomerHandler) Update(c *gin.Context) {
	customerId := c.Param("id")
	n, err := strconv.Atoi(customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err = h.service.Update(n, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

func (h *CustomerHandler) DeleteById(c *gin.Context) {
	customerId := c.Param("id")
	n, err := strconv.Atoi(customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	err = h.service.DeleteById(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
