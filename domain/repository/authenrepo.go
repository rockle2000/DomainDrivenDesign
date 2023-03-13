package repository

import (
	"context"
	"errors"

	"DDD_Project/domain/model/entity"
	"DDD_Project/infrastructure/persistence/datastore"
	"DDD_Project/infrastructure/util"
)

type AuthenticationRepo interface {
	Login(ctx context.Context, customer *entity.Customer) (*entity.Customer, error)
	Register(ctx context.Context, customer *entity.Customer) error
}

type authenticationRepo struct {
	customerStore datastore.CustomerDatastore
}

func NewAuthenticationRepo(cusStore datastore.CustomerDatastore) AuthenticationRepo {
	return &authenticationRepo{
		customerStore: cusStore,
	}
}

func (a *authenticationRepo) Login(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {
	existingCustomer, err := a.customerStore.GetByEmail(ctx, customer.Email)
	if err != nil {
		return nil, errors.New("Invalid credential")
	}
	if !util.CheckPasswordHash(customer.Password, existingCustomer.Password) {
		return nil, errors.New("Invalid credential")
	}
	return existingCustomer, nil
}

func (a *authenticationRepo) Register(ctx context.Context, customer *entity.Customer) error {
	existingCustomer, err := a.customerStore.FindByEmail(ctx, customer.Email)
	if err != nil {
		return err
	}
	if existingCustomer.Id != 0 {
		return errors.New("email had already existed")
	}
	passwordHash, err := util.HashPassword(customer.Password)
	if err != nil {
		return err
	}
	customer.Password = passwordHash
	err = a.customerStore.Create(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}
