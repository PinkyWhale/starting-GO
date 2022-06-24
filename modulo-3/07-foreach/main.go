package main

import "fmt"

func main() {

	animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}

	for indice, elemento := range animales {
		fmt.Println("El indice es: ", indice, ". El valor es: ", elemento)
	}

}
