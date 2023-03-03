package entity

import "DDD_Project/domain/model/valueobject"

type Customer struct {
	Id      string              `json:"id"`
	Name    string              `json:"name"`
	Email   string              `json:"email"`
	Address valueobject.Address `json:"address"`
}
