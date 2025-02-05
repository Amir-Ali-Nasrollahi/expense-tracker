package main

import (
	"fmt"
	"os"
)

func getSwitches() []string {
	arguments := os.Args
	return arguments
}

func getDescription(entry []string) string {
	var descriptionKey int
	for key, value := range entry {
		if value == "--description" {
			descriptionKey = key + 1
		}
	}
	if descriptionKey == 0 {
		return "-"
	}
	return entry[descriptionKey]
}

func getAmount(entry []string) string {
	var amountKey int
	for key, value := range entry {
		if value == "--amount" {
			amountKey = key + 1
		}
	}
	if amountKey == 0 {
		return "-"
	}
	return entry[amountKey]
}

func getId(entry []string) string {
	var Id int
	for key, value := range entry {
		if value == "--id" {
			Id = key + 1
		}
	}
	if Id == 0 {
		return "-"
	}
	return entry[Id]
}

func getCommand(entry []string) (string, error) {
	for _, value := range entry {

		switch value {
		case "add":
			return value, nil
		case "list":
			return value, nil
		case "delete":
			return value, nil
		case "summery":
			return value, nil
		}
	}

	return "", fmt.Errorf("invalid command") 
}

func main() {

	ex := Expense{}

	switches := getSwitches()
	command, err := getCommand(switches)

	if err != nil {
		fmt.Println("please enter the valid command")
		return
	}

	switch command {
	case "add":
		addValue := ex.Add(getAmount(switches), getDescription(switches))
		fmt.Println(addValue)
	case "list":
		listValue := ex.Show()
		fmt.Println(listValue)
	case "summery":
	case "delete":
		deleteValue := ex.Delete(getId(switches))	
		fmt.Println(deleteValue)
	}

}
