package repos

import (
	"database/sql"

	"github.com/wsaefulloh/rest-api-go/models"
)

type initRepo struct {
	db *sql.DB
}

func NewUsers(dbms *sql.DB) *initRepo {
	return &initRepo{dbms}
}

func (r *initRepo) FindAll() (*models.Users, error) {
	query := `SELECT * FROM public.users`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Users
	var users models.User

	for rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.UserName, &users.Email, &users.Role, &users.Password, &users.CreatedAt, &users.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, users)
	}

	return &data, nil
}

func (r *initRepo) Save(user *models.User) error {
	query := `INSERT INTO public.users("name", username, email, "role", "password", created_at, update_at) VALUES($1, $2, $3, $4, $5, $6, $7)`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(user.Name, user.UserName, user.Email, user.Role, user.Password, user.CreatedAt, user.UpdateAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepo) GetPass(username string) (string, error) {
	query := `SELECT "password" FROM public.users WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var password string

	for rows.Next() {
		err := rows.Scan(&password)
		if err != nil {
			return "", err
		}
	}

	return password, nil
}

func (r *initRepo) GetEmail(username string) (string, error) {
	query := `SELECT "email" FROM public.users WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var email string

	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			return "", err
		}
	}

	return email, nil
}

func (r *initRepo) GetUsername(username string) (string, error) {
	query := `SELECT "username" FROM public.users WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var user string

	for rows.Next() {
		err := rows.Scan(&user)
		if err != nil {
			return "", err
		}
	}

	return user, nil
}

func (r *initRepo) GetRole(username string) (string, error) {
	query := `SELECT "role" FROM public.users WHERE "username"=$1`

	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var role string

	for rows.Next() {
		err := rows.Scan(&role)
		if err != nil {
			return "", err
		}
	}

	return role, nil
}
