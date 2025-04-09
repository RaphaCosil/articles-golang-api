package repository

import (
	"github.com/raphacosil/go-study/internal/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) FindAll() ([]model.Customer, error) {
	var customer []model.Customer
	result := r.db.Find(&customer)
	return customer, result.Error
}
