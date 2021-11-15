package models

import "time"

type Product struct {
	Id        int
	Name      string
	Price     int
	Category  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Products []Product

func CreateProduct() *Product {
	return &Product{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
