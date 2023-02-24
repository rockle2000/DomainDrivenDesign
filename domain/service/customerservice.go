package service

import (
	"context"

	"DDD_Project/domain/model/entity"
	"DDD_Project/domain/repository"
)

type CustomerService interface {
	GetCustomer(id string) (*entity.Customer, error)
	GetAllCustomer() ([]*entity.Customer, error)
	CreateCustomer(customer *entity.Customer) error
	UpdateCustomer(customer *entity.Customer) error
	DeleteCustomer(id string) error
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: repo,
	}
}

func (s *customerService) GetCustomer(id string) (*entity.Customer, error) {
	ctx := context.Background()
	return s.customerRepository.FindById(ctx, id)
}

func (s *customerService) CreateCustomer(customer *entity.Customer) error {
	ctx := context.Background()
	return s.customerRepository.Create(ctx, customer)
}
func (s *customerService) UpdateCustomer(customer *entity.Customer) error {
	ctx := context.Background()
	return s.customerRepository.Update(ctx, customer)
}

func (s *customerService) DeleteCustomer(id string) error {
	ctx := context.Background()
	return s.customerRepository.Delete(ctx, id)
}

func (s *customerService) GetAllCustomer() ([]*entity.Customer, error) {
	ctx := context.Background()
	return s.customerRepository.FindAll(ctx)
}
