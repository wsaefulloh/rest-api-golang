package repos

import (
	"database/sql"
	"fmt"

	"github.com/wsaefulloh/rest-api-go/models"
)

type initRepoCategory struct {
	db *sql.DB
}

// constructer
func NewCategory(dbms *sql.DB) *initRepoCategory {
	return &initRepoCategory{dbms}
}

func (r *initRepoCategory) FindAll() (*models.Categories, error) {
	query := `SELECT * FROM public.categories`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Categories
	var cate models.Category
	var cateId int
	var cateName string

	for rows.Next() {
		err := rows.Scan(&cateId, &cateName, &cate.CreatedAt, &cate.UpdateAt)
		if err != nil {
			return nil, err
		}
		// errs := rows.Scan(&cate.Id, &cate.Name, &cate.CreatedAt, &cate.UpdateAt)
		fmt.Println(cateId)
		cate.Id = cateId
		cate.Name = cateName
		data = append(data, cate)
	}

	return &data, nil
}

func (r *initRepoCategory) Save(cate *models.Category) error {
	query := `INSERT INTO public.categories("name", created_at, update_at) VALUES($1, $2, $3)`
	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(cate.Name, cate.CreatedAt, cate.UpdateAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepoCategory) Remove(id string) error {
	query := `DELETE FROM public.categories WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepoCategory) Edit(cate *models.Category, id string) error {
	query := `UPDATE public.categories SET "name" = $1, update_at = $2  WHERE id = $3`
	_, err := r.db.Exec(query, cate.Name, cate.UpdateAt, id)

	if err != nil {
		return err
	}

	return nil
}

// func (r *initRepoCategory) GetCategory(id int) bool {
// 	query := `SELECT "name" FROM public.categories WHERE "id"=$1`

// 	rows, err := r.db.Query(query, id)

// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	defer rows.Close()

// 	var name string

// 	for rows.Next() {
// 		err := rows.Scan(&name)
// 		if err != nil {
// 			fmt.Println(err)
// 			return false
// 		}
// 	}

// 	return len([]rune(name)) != 0
// }
