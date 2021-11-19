package repos

import (
	"database/sql"

	"github.com/wsaefulloh/rest-api-go/models"
)

type initRepoProduct struct {
	db *sql.DB
}

func NewProduct(dbms *sql.DB) *initRepoProduct {
	return &initRepoProduct{dbms}
}

func (r *initRepoProduct) FindAll() (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}

func (r *initRepoProduct) Save(prod *models.Product) error {
	query := `INSERT INTO public.products("name", price, category, created_at, update_at) VALUES($1, $2, $3, $4, $5)`
	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(prod.Name, prod.Price, prod.Category, prod.CreatedAt, prod.UpdateAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepoProduct) Remove(id string) error {
	query := `DELETE FROM public.products WHERE id = $1`
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepoProduct) Edit(prod *models.Product, id string) error {
	query := `UPDATE public.products SET "name" = $1, price = $2, category = $3, update_at = $4  WHERE id = $5`
	_, err := r.db.Exec(query, prod.Name, prod.Price, prod.Category, prod.UpdateAt, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepoProduct) SearchProductName(name string) (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id WHERE public.products.name ILIKE '%' || $1 || '%' ORDER BY update_at DESC`
	rows, err := r.db.Query(query, name)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}

func (r *initRepoProduct) SearchProductCategory(name string) (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id WHERE public.categories.name ILIKE '%' || $1 || '%' ORDER BY update_at DESC`
	rows, err := r.db.Query(query, name)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}

func (r *initRepoProduct) FindbyCategory() (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id ORDER BY public.categories.name ASC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}

func (r *initRepoProduct) FindbyDateDESC() (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id ORDER BY update_at DESC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}

func (r *initRepoProduct) FindbyDateASC() (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id ORDER BY update_at ASC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}

func (r *initRepoProduct) FindbyPriceDESC() (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id ORDER BY CAST(public.products.price AS DOUBLE PRECISION) DESC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}

func (r *initRepoProduct) FindbyPriceASC() (*models.Products, error) {
	query := `SELECT public.products.id, public.products.name, public.products.price, public.categories.name,public.products.created_at, public.products.update_at FROM public.products INNER JOIN public.categories ON public.products.category = public.categories.id ORDER BY CAST(public.products.price AS DOUBLE PRECISION) ASC`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Products
	var prod models.Product

	for rows.Next() {
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Price, &prod.Category, &prod.CreatedAt, &prod.UpdateAt)
		if err != nil {
			return nil, err
		}
		data = append(data, prod)
	}

	return &data, nil
}
