package entity

import (
	"time"

	"DDD_Project/domain/model/valueobject"
)

type Order struct {
	OrderBy         Customer
	Products        []Product
	Total           int
	OrderDate       *time.Time
	DeliveryAddress string
	Note            string
	Payment         valueobject.Payment
}
