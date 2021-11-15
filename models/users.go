package models

import "time"

type User struct {
	Id        int
	Name      string
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Users []User

func CreateUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
