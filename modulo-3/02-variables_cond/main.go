package main

import "fmt"

func main() {
	//creadion de cariable en if y utilizacion
	if nombre, edad := "Cody", 7; nombre == "Cody" {
		fmt.Println("Hola", nombre, "te damos la bienvenida al curso de GO!")
	} else {
		fmt.Println("Los valores son: ", nombre, edad)
	}
}
