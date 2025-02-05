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
		test := ex.Add(getAmount(switches), getDescription(switches))
		fmt.Println(test)
	case "list":
	case "summery":
	case "delete":
	}

}
