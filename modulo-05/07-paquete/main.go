package main

import (
	"fmt"

	"./Pack" // paquete publico porque esta en Mayuscula
)

func main() {

	curso := Pack.Curso{Titulo: "Curso profesional de GO"}

	fmt.Println(curso.GetTitulo)
}
