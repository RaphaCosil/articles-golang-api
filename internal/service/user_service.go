package service

import(
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/repository"
)

type UserService interface {
	FindAll() ([]model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) FindAll() ([]model.User, error) {
	return s.repo.FindAll()
}