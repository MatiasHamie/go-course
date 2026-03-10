package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals:
// 1) Validate the user input
// 		=> Show error message & exit if invalid input is provided
// 		- No negative numbers
// 		- Not 0
// 2) Store calculated results into file

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.2f", ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(inputText string) (float64, error) {
	var userInputValue float64
	fmt.Print(inputText)
	fmt.Scan(&userInputValue)

	if userInputValue <= 0 {
		return 0, errors.New("Value must be a positive number")
	}

	return userInputValue, nil
}

func calculateValues(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
