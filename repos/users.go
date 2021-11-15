package repos

import (
	"database/sql"

	"github.com/wsaefulloh/rest-api-go/configs/db"
	"github.com/wsaefulloh/rest-api-go/models"
)

type InitRepo struct {
	db *sql.DB
}

func New() *InitRepo {
	db, _ := db.New()
	return &InitRepo{db}
}

func (r *InitRepo) FindAll() (*models.Users, error) {
	query := `SELECT * FROM public.users`
	rows, err := r.db.Query(query)

	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var data models.Users
	var users models.User

	for rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.UserName, &users.Email, &users.Password, &users.CreatedAt, &users.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, users)
	}

	return &data, nil
}

func (r *InitRepo) Save(user *models.User) error {
	query := `INSERT INTO public.users("name", username, email, "password", created_at, update_at) VALUES($1, $2, $3, $4, $5, $6)`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(user.Name, user.UserName, user.Email, user.Password, user.CreatedAt, user.UpdateAt)

	if err != nil {
		return err
	}

	return nil
}
