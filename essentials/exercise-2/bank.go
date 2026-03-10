package main

// si se quiere instalar dependencias, como en el FE con un manejador de paquetes, usas el comando
// ❯ go get + url del paquete que queres
// por ej ❯ go get github.com/Pallinder/go-randomdata
import (
	"fmt"
	// para usar un package custom, tenemos q poner la url del go.mod+el nombre del package
	// para usar las funciones de ese package, esas funciones tienen que empezar con mayuscula
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")

		// panic("Can't continue")
		// el panic lo que hace es crashear la app
		// lo que imprimiria seria algo asi:
		// 		Failed to find balance file.
		// ----------
		// panic: Can't continue

		// goroutine 1 [running]:
		// main.main()
		//         /Users/sarasa/Desktop/Courses/go/lesson-1/exercise-2/bank.go:65 +0x62c
		// exit status 2

	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	// for i := 0; i < 2; i++ {
	for { // ya con hacer esto tengo un loop infinito
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				// return
				continue //le dice a go que skipee el resto y continue con la siguiente iteracion
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("How much: ")

			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount, you can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawAmount

			fmt.Println("Balance updated! New Amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our Bank")
			return

			// break // no se puede usar break dentro de un switch, go lo hace automaticamente, pero en una interacion con puros ifs si
		}
	}

}
