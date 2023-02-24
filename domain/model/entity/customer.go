package entity

import "DDD_Project/domain/model/valueobject"

type Customer struct {
	Id      string
	Name    string
	Email   string
	Address valueobject.Address
}
