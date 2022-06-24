package main

import "fmt"

type Operacion func(balance, cantidad int) int

func suma(num1, num2 int) int { //Operacion
	return num1 + num2
}

func resta(num1, num2 int) int { // Operacion
	return num1 - num2
}

func multiplicacion(num1, num2 int) int {
	return num1 * num2
}

func ejecutarOperacion(funcion Operacion) {
	fmt.Println("Antes de la Operacion")

	resultado := funcion(100, 50)
	fmt.Println("El resultado es:", resultado)
	fmt.Println("Despues de la operacion")

}

func main() {
	ejecutarOperacion(multiplicacion)
}
