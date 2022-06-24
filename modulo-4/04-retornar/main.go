package main

import "fmt"

type Operacion func(balance, cantidad int) int

func crearOperacion(tipo string) Operacion {

	if tipo == "suma" {
		//retornamos una funcion anonima
		return func(balance, cantidad int) int { return balance + cantidad }
	} else if tipo == "resta" {
		return func(balance, cantidad int) int { return balance - cantidad }
	} else {
		return func(balance, cantidad int) int { return balance * cantidad }
	}

}

func main() {
	suma := crearOperacion("suma")
	resultado := suma(40, 50)

	fmt.Println("El rrsultado es:", resultado)
}
