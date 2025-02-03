package main

import (
	"fmt"
	"os"
)

func getSwitches() []string {
	arguments := os.Args
	return arguments[1:]
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

func main() {

	fmt.Println(getDescription(getSwitches()), getAmount(getSwitches()))

}
