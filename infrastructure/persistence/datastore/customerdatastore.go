package datastore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"DDD_Project/domain/model/entity"
)

type CustomerDatastore interface {
	GetById(ctx context.Context, id int) (*entity.Customer, error)
	GetAll(ctx context.Context) ([]*entity.Customer, error)
	Create(ctx context.Context, customer *entity.Customer) error
	Update(ctx context.Context, customer *entity.Customer) error
	Delete(ctx context.Context, id int) error
	GetByEmail(ctx context.Context, email string) (*entity.Customer, error)
	FindByEmail(ctx context.Context, email string) (*entity.Customer, error)
}

type customerDatastore struct {
	DB *sqlx.DB
}

func NewCustomerDatastore(db *sqlx.DB) CustomerDatastore {
	return &customerDatastore{
		DB: db,
	}
}

func (c *customerDatastore) GetById(ctx context.Context, id int) (*entity.Customer, error) {
	var customer entity.Customer
	sqlQuery := fmt.Sprint("SELECT c.id,c.name,c.email,a.id as \"address.id\",a.city as \"address.city\", c.password " +
		"FROM customer c JOIN address a ON c.address_id = a.id WHERE c.id = $1")
	err := c.DB.GetContext(ctx, &customer, sqlQuery, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customerDatastore) GetAll(ctx context.Context) ([]*entity.Customer, error) {
	var listCustomer []*entity.Customer
	sqlQuery := fmt.Sprintf("SELECT c.id,c.name,c.email, a.id as \"address.id\",a.city as \"address.city\", c.password FROM customer c JOIN address a ON c.address_id = a.id")
	err := c.DB.SelectContext(ctx, &listCustomer, sqlQuery)

	if err != nil {
		return nil, err
	}
	return listCustomer, nil
}

func (c *customerDatastore) Create(ctx context.Context, customer *entity.Customer) error {
	sqlInsert := fmt.Sprintf("INSERT INTO customer (name, email, password) VALUES ($1,$2,$3)")
	_, err := c.DB.ExecContext(ctx, sqlInsert, customer.Name, customer.Email, customer.Password)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerDatastore) Delete(ctx context.Context, id int) error {
	sqlDelete := fmt.Sprintf("DELETE customer WHERE id = $1")
	_, err := c.DB.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		return err
	}
	return nil
}
func (c *customerDatastore) Update(ctx context.Context, customer *entity.Customer) error {
	sqlUpdate := fmt.Sprintf("UPDATE customer SET name = $1 email = $2 WHERE id = $3")
	_, err := c.DB.ExecContext(ctx, sqlUpdate, customer.Name, customer.Email, customer.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerDatastore) GetByEmail(ctx context.Context, email string) (*entity.Customer, error) {
	var customer entity.Customer
	sqlQuery := fmt.Sprintf("SELECT c.id, c.name, c.email, c.password FROM customer c WHERE c.email = $1")
	err := c.DB.GetContext(ctx, &customer, sqlQuery, email)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customerDatastore) FindByEmail(ctx context.Context, email string) (*entity.Customer, error) {
	var customer entity.Customer
	sqlQuery := fmt.Sprintf("SELECT c.id, c.name, c.email, c.password FROM customer c WHERE c.email = $1")
	err := c.DB.GetContext(ctx, &customer, sqlQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return &customer, nil
		}
		return nil, err
	}
	return &customer, nil
}
