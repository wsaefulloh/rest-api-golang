package models

import (
	"strconv"
	"time"
)

type History struct {
	Id      int
	Invoice string
	Cashier string
	Date    string
	Order   string
	Count   int
	Total   int
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
