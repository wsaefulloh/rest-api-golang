package models

import "time"

type Category struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Categories []Category

func CreateCategory() *Category {
	return &Category{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
