package main

import (
	"errors"
	"fmt"
)

func divisor(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir sobre 0.")
	} else {
		return dividendo / divisor, nil // nil representa la ausencia de valor
	}
}

func main() {

	if resultado, err := divisor(10, 2); err != nil {
		//fmt.Println(err)
		panic(err) // condicionamos sobre un error
	} else {
		fmt.Println("El resultado es:", resultado)
	}

}
