package main

import (
	"fmt"
	"reflect"
)

// go build main.go -> un arhivo ejecutable
// go run main.go
const (
	Domingo int = iota + 1 //0
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	//-------- creacion de variables
	var nombre string
	var edad int
	var altura = 1.65
	curso := "curso de GO"

	nombre = "Erica"
	edad = 31

	// nombre := "Erica"
	// edad := 31

	fmt.Println(nombre, edad, altura)
	fmt.Println(curso)

	//-------- operadores relacionales
	/*
		>
		<
		>=
		<=
		==
		!=
	*/

	var num1 = 5
	var resultado = num1 >= 5 // Bool

	fmt.Println(resultado)

	//-------- operadores logicos
	/*
		&& and
		|| or
		! not
	*/

	resultadoA := 20 == 20 && 30 == 30 && 20 > 40

	fmt.Println(resultadoA)

	// -------- Secuencia de Valores

	fmt.Println(Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)

	//-------- acciones con  string
	cursar := "Curso Profesional de GO"

	longitud := len(cursar)     // int
	primerCaracter := cursar[0] // char
	ultimoCaracter := curso[len(curso)-1]

	fmt.Println("La longitud del texto es:", longitud)
	fmt.Println("El primer caracter en la letra de posicion:", primerCaracter)
	fmt.Printf("%c\n", primerCaracter)          // se imprimer l primer caracter
	fmt.Printf("%c\n", ultimoCaracter)          // se imprime el ultimo caracter
	fmt.Println(reflect.TypeOf(primerCaracter)) // uint8 -> un entero positivo de 8 bits

	//-------- Lectura de caracteres

	var nombre_ string
	var edad_ int
	var altura_ float32

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre_)

	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad_)
	/*
		fmt.Print("Ingresa tu altura: ")
		fmt.Scanf("%f", &altura_)
	*/

	fmt.Printf("Hola %s con una edad %d y una altura de %.2f\n", nombre_, edad_, altura_)
}
