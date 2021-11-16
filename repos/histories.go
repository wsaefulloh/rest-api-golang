package repos

import (
	"database/sql"

	"github.com/wsaefulloh/rest-api-go/configs/db"
	"github.com/wsaefulloh/rest-api-go/models"
)

type InitRepoHistory struct {
	db *sql.DB
}

func NewHistory() *InitRepoHistory {
	db, _ := db.New()
	return &InitRepoHistory{db}
}

func (r *InitRepoHistory) FindAll() (*models.Histories, error) {
	query := `SELECT public.histories.id, public.histories.invoice, public.histories.cashier, public.histories.date, public.products.name, public.histories.count, public.products.price FROM public.histories INNER JOIN public.products ON public.histories.order = public.products.id`
	rows, err := r.db.Query(query)

	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var data models.Histories
	var histo models.History
	var histoPrice int

	for rows.Next() {
		err := rows.Scan(&histo.Id, &histo.Invoice, &histo.Cashier, &histo.Date, &histo.Order, &histo.Count, &histoPrice)
		if err != nil {
			return nil, err
		}
		histo.Total = histo.Count * histoPrice
		data = append(data, histo)
	}

	return &data, nil
}

func (r *InitRepoHistory) Save(histo *models.History) error {
	query := `INSERT INTO public.histories("invoice", "cashier", "date", "order", "count") VALUES($1, $2, $3, $4, $5)`
	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(histo.Invoice, histo.Cashier, histo.Date, histo.Order, histo.Count)

	if err != nil {
		return err
	}

	return nil
}

func (r *InitRepoHistory) Remove(id string) error {
	query := `DELETE FROM public.histories WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *InitRepoHistory) Edit(histo *models.History, id string) error {
	query := `UPDATE public.histories SET "invoice" = $1, "cashier" = $2, "date" = $3, "order" = $4, "count" = $5  WHERE id = $6`
	_, err := r.db.Exec(query, histo.Invoice, histo.Cashier, histo.Date, histo.Order, histo.Count, id)

	if err != nil {
		return err
	}

	return nil
}
