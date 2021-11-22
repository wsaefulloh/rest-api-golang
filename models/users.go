package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Users []User

func CreateUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
