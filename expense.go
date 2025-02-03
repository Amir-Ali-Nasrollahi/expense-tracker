package main

import "time"

type Expense struct {
	ID int
	Date string
	Description string
	Amount string
}

func (e *Expense) Add(amount string, desc string) {
	e.Amount = amount
	e.Description = desc
	e.Date = time.Now().Format("Y-M-D")
}

