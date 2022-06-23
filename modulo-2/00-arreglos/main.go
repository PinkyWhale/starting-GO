package main

import "fmt"

func main() {

	//----- Creacion de arreglos
	// Creacion comun
	var numeros [5]int

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300
	numeros[3] = 400
	numeros[4] = 500

	fmt.Println("El arreglo es el siguiente: ", numeros)
	fmt.Println(numeros[2])

	// Creacion sin determinar tipo de arreglo y cantidad
	nums := [...]int{100, 200, 300, 400, 500}

	fmt.Println(nums)

	// Practica

	monedas := [...]string{0: "Peso Mex", 2: "Peso Arg", 1: "Dolar"}

	fmt.Println(monedas)

	//Sub areglo
	subMonedas := monedas[1:2] //slice

	fmt.Println(subMonedas)
}
