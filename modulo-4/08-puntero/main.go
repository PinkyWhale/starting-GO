package main

import "fmt"

func modificarValor(parametro *string) {
	*parametro = "Cambio de valor" //se toma la referencia y se la modifica

}

func main() {
	var curso = "Curso profesional de GO!"

	fmt.Println("Antes de la funcion:", curso)

	modificarValor(&curso) //Referencia a la variable

	fmt.Println("Despues de la funcion", curso)
}
