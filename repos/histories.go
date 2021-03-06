package repos

import (
	"database/sql"

	"github.com/wsaefulloh/rest-api-go/models"
)

type initRepoHistory struct {
	db *sql.DB
}

func NewHistory(dbms *sql.DB) *initRepoHistory {
	return &initRepoHistory{dbms}
}

func (r *initRepoHistory) FindAll() (*models.Histories, error) {
	query := `SELECT public.histories.id, public.histories.invoice, public.histories.cashier, public.histories.date, public.products.name, public.histories.count, public.products.price FROM public.histories INNER JOIN public.products ON public.histories.order = public.products.id`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

func (r *initRepoHistory) Save(histo *models.History) error {
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

func (r *initRepoHistory) Remove(id string) error {
	query := `DELETE FROM public.histories WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepoHistory) Edit(histo *models.History, id string) error {
	query := `UPDATE public.histories SET "invoice" = $1, "cashier" = $2, "date" = $3, "order" = $4, "count" = $5  WHERE id = $6`
	_, err := r.db.Exec(query, histo.Invoice, histo.Cashier, histo.Date, histo.Order, histo.Count, id)

	if err != nil {
		return err
	}

	return nil
}
