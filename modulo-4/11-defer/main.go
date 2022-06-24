package main

import "fmt"

func funcion1() {
	fmt.Println("Hola soy la primera funcion!")

}
func funcion2() {
	fmt.Println("Hola soy la segunda funcion!")

}

func main() {
	defer funcion1() // la funcio se va a ejecutar cuando main finalice
	funcion2()

}
