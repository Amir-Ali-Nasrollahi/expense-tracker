package main

import (
	"fmt"
	"os"
	"regexp"
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
		formatted = fmt.Sprintf("ID: %v, Description: %v, Date: %v, Amount: %v$\n", e.ID, e.Description, e.Date, e.Amount)
		os.WriteFile("./store.txt", []byte(formatted), 0644)

		return "# Expense added successfully (ID:" + strconv.Itoa(e.ID) + ")"
	}

	splited := strings.Split(string(file), "\n")
	e.ID = len(splited)

	formatted = fmt.Sprintf("%v ID: %v, Description: %v, Date: %v, Amount: %v$\n", string(file), e.ID, e.Description, e.Date, e.Amount)
	os.WriteFile("./store.txt", []byte(formatted), 0644)

	return "# Expense added successfully (ID:" + strconv.Itoa(e.ID) + ")"
}

func (e *Expense) Delete(id string) string {

	file, err := os.ReadFile("./store.txt")
	if err != nil {
		return "you should first add expense then delete that"
	}
	storeExpense := strings.Split(string(file), "\n")

	os.Remove("./store.txt")

	WriteFile, _ := os.OpenFile("./store.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	for key, value := range storeExpense[:len(storeExpense)-1] {
		storeId := strings.Split(value, ",")
		reg, _ := regexp.Compile(id)

		fmt.Println(reg.MatchString(storeId[0]), key)
		if !reg.MatchString(storeId[0]) {
			WriteFile.WriteString(value + "\n")
		}
	}

	return "item deleted successfuly"
}

func (e Expense) Show() string {

	file , err := os.ReadFile("./store.txt")
	if err != nil {
		return "First, you need to add expenses"
	}
	return string(file)
}

func (e Expense) Summary(SpecificMonth int) {
	// nothing again
	// nothing just testing git
}
