package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	ID          int
	Date        string
	Description string
	Amount      string
}

func (e *Expense) Add(amount string, desc string) string {
	var formatted string

	e.Amount = amount
	e.Description = desc
	e.Date = fmt.Sprintf("%v-%v-%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day())

	file, err := os.ReadFile("./store.txt")
	if err != nil {
		e.ID = 1
		formatted = fmt.Sprintf("ID: %v , Description: %v, Date: %v, Amount: %v$\n", e.ID, e.Description, e.Date, e.Amount)
		os.WriteFile("./store.txt", []byte(formatted), 0644)

		return "# Expense added successfully (ID:" + strconv.Itoa(e.ID) + ")"
	}

	splited := strings.Split(string(file), "\n")
	e.ID = len(splited)

	formatted = fmt.Sprintf("%v ID: %v , Description: %v, Date: %v, Amount: %v$\n", string(file), e.ID, e.Description, e.Date, e.Amount)
	os.WriteFile("./store.txt", []byte(formatted), 0644)

	return "# Expense added successfully (ID:" + strconv.Itoa(e.ID) + ")"
}
