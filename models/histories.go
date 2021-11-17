package models

import (
	"strconv"
	"time"
)

type History struct {
	Id      int    `json:"id"`
	Invoice string `json:"invoice"`
	Cashier string `json:"cashier"`
	Date    string `json:"date"`
	Order   string `json:"order"`
	Count   int    `json:"count"`
	Total   int    `json:"total"`
}

type Histories []History

func CreateHistory() *History {
	current_time := time.Now()
	day := strconv.Itoa(current_time.Day())
	month := strconv.Itoa(int(current_time.Month()))
	year := strconv.Itoa(current_time.Year())
	t := day + "-" + month + "-" + year
	return &History{
		Date: t,
	}
}
