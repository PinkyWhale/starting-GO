package main

import "fmt"

func main() {

	// Bloque es una coleccion de sentencias agrupadas por bloques de llaves

	var curso = "Curso profesional de GO!"
	var version = 10

	{ //Bloque 2

		fmt.Println(curso)

		{ // Bloque3
			var version = 5
			fmt.Println("Estamos en el bloque 3", version)
			fmt.Println(curso)
		}

		fmt.Println("Estamos en el bloque 2", version)

	}

	fmt.Println(curso)

}
