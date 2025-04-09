package service

import(
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/repository"
)

type CustomerService interface {
	FindAll() ([]model.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(r repository.CustomerRepository) CustomerService {
	return &customerService{r}
}

func (s *customerService) FindAll() ([]model.Customer, error) {
	return s.repo.FindAll()
}