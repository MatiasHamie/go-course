/*
Cuando queremos ejecutar el codigo, se puede usar el go run nombre_archivo.go
o, si tenemos el modulo, se hace go run . y ya esta
*/
package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturned := 5.5

	showOutputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	showOutputText("Expected return rate: ")
	fmt.Scan(&expectedReturned)

	showOutputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturned, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func showOutputText(text string) {
	fmt.Print(text)
}

// Si retorno un solo valor, no hace falta poner el return type entre parentesis
// func calculateFutureValue(investmentAmount, expectedReturned, years float64) float64 {
// func calculateFutureValue(investmentAmount, expectedReturned, years float64) (float64, float64) {
func calculateFutureValue(investmentAmount, expectedReturned, years float64) (fv float64, rfv float64) {
	// fv := investmentAmount * math.Pow(1+expectedReturned/100, years)
	// no uso el := aca porque en el return type ya cree las variables -> (fv float64, rfv float64)
	fv = investmentAmount * math.Pow(1+expectedReturned/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv // se pueden retornar varios valores solo al separar con comas
	// si pongo return solo es valido, ya automaticamente go sabe que tiene q retornar (fv float64, rfv float64) estas variables
	// return
}

func aprendiendoCuartoMain() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturned := 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturned)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturned/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// diferentes formas de hacer print
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %v\n", futureValue)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f", futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func aprendiendoTercerMain() {
	// tambien se puede hacer esto, si son del mismo tipo
	// investmentAmount, years, expectedReturned := 1000.0, 10.0, 5.5
	// o esto, especificando el type
	// 	investmentAmount, years, expectedReturned float64 = 1000.0, 10.0, 5.5

	investmentAmount := 1000.0
	years := 10.0
	expectedReturned := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturned/100, years)

	fmt.Println(futureValue)
}

func aprendiendoSegundoMain() {
	var investmentAmount float64 = 1000
	// Si estoy contento con el tipo inferido por go, cuando quiero crear y asignar una variable
	// uso el operador := el cual no hace falta poner el type de forma explicita
	// es mas, si pongo el type me tira error
	expectedReturned := 5.5
	var years float64 = 10

	futureValue := investmentAmount * math.Pow(1+expectedReturned/100, years)

	fmt.Println(futureValue)
}

func aprendiendoPrimerMain() {
	// fmt.Print("Hello World")
	// Si le ponemos el tipo explicito a la variable, se trata como float de una
	// no hace falta parsear nada con utils
	var investmentAmount float64 = 1000
	var expectedReturned = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturned/100, years)

	fmt.Println(futureValue)
}
