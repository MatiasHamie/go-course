package fileops // un paquete custom tiene que tener su folder que se llame igual

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// para que pueda ser usada una funcion desde otro package, tiene q empezar con mayuscula
func GetFloatFromFile(fileName string) (float64, error) {
	// si readFile tiene un error, por ej si el archivo no existe, no crashea la app, te devuelve un error y un valor 0
	// esto seria data, error, pero en go si usas el _, le decis a go q no lo vas a usar ahora
	// data, _ := os.ReadFile(accountBalanceFile)
	data, err := os.ReadFile(fileName)

	// nil = la ausencia de un valor utilizable
	// error en Go no es “un error”, es un valor
	// En Go, error es un tipo de dato, no una excepción.
	// 	result, err := divide(10, 2)

	// if err != nil {
	//     fmt.Println("hubo error:", err)
	// }

	// Caso correcto:
	// result = 5
	// err = nil
	// la operación fue exitosa

	// result = 0
	// err = "division by zero"
	// la operación falló
	// err == nil   → todo OK
	// err != nil   → algo falló
	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	valueText := string(data) // string() funciona con los []byte
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	// []byte(valueText), el []byte transforma a bytes el texto que yo quiero
	// los files premissions 0644 = https://www.redhat.com/en/blog/linux-file-permissions-explained
	os.WriteFile(fileName, []byte(valueText), 0644)
}
