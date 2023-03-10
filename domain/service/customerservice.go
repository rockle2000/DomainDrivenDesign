package service

import (
	"context"

	"DDD_Project/domain/model/entity"
	"DDD_Project/domain/repository"
)

type CustomerReq struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerService interface {
	GetCustomer(id int) (*entity.Customer, error)
	GetAllCustomer() ([]*entity.Customer, error)
	CreateCustomer(customer CustomerReq) error
	UpdateCustomer(customer CustomerReq) error
	DeleteCustomer(id int) error
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: repo,
	}
}

func (s *customerService) GetCustomer(id int) (*entity.Customer, error) {
	ctx := context.Background()
	return s.customerRepository.FindById(ctx, id)
}

func (s *customerService) CreateCustomer(customer CustomerReq) error {
	ctx := context.Background()
	customerEntity := &entity.Customer{
		Id:       customer.Id,
		Name:     customer.Name,
		Email:    customer.Email,
		Password: customer.Password,
	}
	return s.customerRepository.Create(ctx, customerEntity)
}
func (s *customerService) UpdateCustomer(customer CustomerReq) error {
	ctx := context.Background()
	customerEntity := &entity.Customer{
		Id:       customer.Id,
		Name:     customer.Name,
		Email:    customer.Email,
		Password: customer.Password,
	}
	return s.customerRepository.Update(ctx, customerEntity)
}

func (s *customerService) DeleteCustomer(id int) error {
	ctx := context.Background()
	return s.customerRepository.Delete(ctx, id)
}

func (s *customerService) GetAllCustomer() ([]*entity.Customer, error) {
	ctx := context.Background()
	return s.customerRepository.FindAll(ctx)
}
