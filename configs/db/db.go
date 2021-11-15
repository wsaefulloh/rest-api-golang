package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	host := "localhost"
	user := "devops"
	password := "abcd1234"
	dbName := "Golang"

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	db, err := connect(config)

	if err != nil {
		return nil, err
	}

	fmt.Println("Database ready")

	return db, nil
}

func connect(config string) (*sql.DB, error) {
	db, err := sql.Open("postgres", config)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
