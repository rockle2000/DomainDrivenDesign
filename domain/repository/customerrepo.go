package repository

import (
	"context"

	"DDD_Project/domain/model/entity"
	"DDD_Project/infrastructure/persistence/datastore"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *entity.Customer) error
	FindById(ctx context.Context, id string) (*entity.Customer, error)
	FindAll(ctx context.Context) ([]*entity.Customer, error)
	Update(ctx context.Context, customer *entity.Customer) error
	Delete(ctx context.Context, id string) error
}

type customerRepository struct {
	dataStore datastore.CustomerDatastore
}

func NewCustomerRepository(dataStore datastore.CustomerDatastore) CustomerRepository {
	return &customerRepository{dataStore: dataStore}
}

func (cr *customerRepository) Create(ctx context.Context, customer *entity.Customer) error {
	return cr.dataStore.Create(ctx, customer)
}

func (cr *customerRepository) FindById(ctx context.Context, id string) (*entity.Customer, error) {
	return cr.dataStore.GetById(ctx, id)
}

func (cr *customerRepository) FindAll(ctx context.Context) ([]*entity.Customer, error) {
	return cr.dataStore.GetAll(ctx)
}

func (cr *customerRepository) Update(ctx context.Context, customer *entity.Customer) error {
	return cr.dataStore.Update(ctx, customer)
}

func (cr *customerRepository) Delete(ctx context.Context, id string) error {
	return cr.dataStore.Delete(ctx, id)
}
