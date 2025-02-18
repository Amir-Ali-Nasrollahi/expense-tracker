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
	if err != nil || len(file) == 0 {
		e.ID = 1
		formatted = fmt.Sprintf("ID: %v, Description: %v, Date: %v, Amount: %v$\n", e.ID, e.Description, e.Date, e.Amount)
		os.WriteFile("./store.txt", []byte(formatted), 0644)

		return "# Expense added successfully (ID:" + strconv.Itoa(e.ID) + ")"
	}

	splited := strings.Split(string(file), "\n")

	NewSplitedByVirgol := strings.Split(splited[len(splited)-2], ",")
	reg, _ := regexp.Compile(`[0-9]+`)

	// find a string in first Item of that ( ID item )
	lastId := reg.FindString(NewSplitedByVirgol[0])
	e.ID, _ = strconv.Atoi(lastId)
	e.ID++

	formatted = fmt.Sprintf("%vID: %v, Description: %v, Date: %v, Amount: %v$\n", string(file), e.ID, e.Description, e.Date, e.Amount)
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

	for _, value := range storeExpense[:len(storeExpense)-1] {
		NewSplitedByVirgol := strings.Split(value, ",")
		reg, _ := regexp.Compile(id)

		if !reg.MatchString(NewSplitedByVirgol[0]) {
			WriteFile.WriteString(value + "\n")
		}
	}

	return "item deleted successfuly"
}

func (e Expense) Show() string {

	file, err := os.ReadFile("./store.txt")
	if err != nil {
		return "First, you need to add expenses"
	}
	return string(file)
}

func (e Expense) Summary(SpecificMonth int) string {

	var sum int
	file, _ := os.ReadFile("./store.txt")
	splited := strings.Split(string(file), "\n")
	if SpecificMonth == 0 {
		for _, value := range splited[:len(splited)-1] {

			NewSplitedByVirgol := strings.Split(value, ",")
			reg, _ := regexp.Compile(`[0-9]+`)

			price := reg.FindString(NewSplitedByVirgol[len(NewSplitedByVirgol)-1])
			summery, _ := strconv.Atoi(price)
			sum += summery
		}
		return "# Total expenses : " + strconv.Itoa(sum) + "$"
	}

	var flag bool = true
	for _, value := range splited[:len(splited)-1] {

		NewSplitedByVirgol := strings.Split(value, ",")
		reg, _ := regexp.Compile(`[0-9]+`)

		Date := reg.FindAllString(NewSplitedByVirgol[len(NewSplitedByVirgol)-2], -1)
		month, _ := strconv.Atoi(Date[1])

		if month == SpecificMonth {
			flag = false
			price := reg.FindString(NewSplitedByVirgol[len(NewSplitedByVirgol)-1])
			summery, _ := strconv.Atoi(price)
			sum += summery
		}
	}

	monthName := time.Month(SpecificMonth).String()

	if flag == true {
		return "# We dont have any expense in " + monthName 
	}
	return "# Total expenses in " + monthName + " : " + strconv.Itoa(sum) + "$"

}
