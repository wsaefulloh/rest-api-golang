package models

import "time"

type Product struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Products []Product

func CreateProduct() *Product {
	return &Product{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
