package service

import(
	"github.com/raphacosil/go-study/internal/model"
	"github.com/raphacosil/go-study/internal/repository"
)

type CustomerService interface {
	FindAll() ([]model.Customer, error)
	FindById(customerId int) (model.Customer, error)
	Create(customer model.Customer) (*model.Customer, *error)
	Update(customerId int, customer model.Customer) error
	DeleteById(customerId int) error
	FindByName(name string) ([]model.Customer, error)
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

func (s *customerService) FindById(customerId int) (model.Customer, error) {
	return s.repo.FindById(customerId)
}

func (s *customerService) Create(customer model.Customer) (*model.Customer, *error) {
	return s.repo.Create(customer)
}

func (s *customerService) Update(customerId int, customer model.Customer) error {
	return s.repo.Update(customerId, customer)
}

func (s *customerService) DeleteById(customerId int) error {
	return s.repo.DeleteById(customerId)
}

func (s *customerService) FindByName(name string) ([]model.Customer, error) {
	return s.repo.FindByName(name)
}
