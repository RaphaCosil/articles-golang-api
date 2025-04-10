package repository

import (
	"github.com/raphacosil/go-study/internal/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
	FindById(customerId int) (model.Customer, error)
	Create(customer model.Customer) (*model.Customer, *error)
	Update(customerId int, customer model.Customer) error
	DeleteById(customerId int) error
	FindByName(name string) ([]model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	result := r.db.Find(&customers)
	return customers, result.Error
}

func (r *customerRepository) FindById(customerId int) (model.Customer, error) {
	var customer model.Customer
	result := r.db.First(&customer, customerId)
	return customer, result.Error
}

func (r *customerRepository) Create(customer model.Customer) (*model.Customer, *error) {
	err := r.db.Create(&customer).Error
	if err !=nil {
		return nil, &err
	}
	return &customer, nil
}

func (r *customerRepository) Update(customerId int, customer model.Customer) error {
	return r.db.Model(&model.Customer{}).Where("customer_id = ?", customerId).Updates(customer).Error
}

func (r *customerRepository) DeleteById(customerId int) error {
	return r.db.Delete(&model.Customer{}, customerId).Error
}

func (r *customerRepository) FindByName(name string) ([]model.Customer, error) {
	var customers []model.Customer
	result := r.db.Where("name ILIKE ?", "%" + name + "%").Find(&customers)
	return customers, result.Error
}
