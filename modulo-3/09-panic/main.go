package main

import "fmt"

func main() {

	// panic function
	var dividendo, divisor int

	fmt.Print("Ingresar un valor para el dividendo: ")
	fmt.Scanf("%d", &dividendo)

	fmt.Print("Ingresar un valor para el divisor: ")
	fmt.Scanf("%d", &divisor)

	if divisor == 0 {
		panic("No es posible dividir sobre 0...") // se valida el error
	}

	resultado := dividendo / divisor

	fmt.Println(resultado)

}
