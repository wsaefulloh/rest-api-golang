package models

import "time"

type Category struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Categories []Category

func CreateCategory() *Category {
	return &Category{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
