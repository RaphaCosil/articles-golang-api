package repository

import (
	"github.com/raphacosil/go-study/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	findAll() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) findAll() ([]model.User, error) {
	var user []model
	result := r.db.Find(&users)
	return users, result.Error
}
