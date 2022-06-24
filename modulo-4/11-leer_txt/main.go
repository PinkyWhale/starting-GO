package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		panic("No fue posible obtener el archivo")
	}

	defer func() {
		fmt.Println("Cerramos el archivo.")
		file.Close()
	}() //aca llamamos a la funcion def

	contenido := make([]byte, 254)

	longitud, _ := file.Read(contenido) // entrega 2 parametos la longitud y el error

	contenidoArchivo := string(contenido[0:longitud])

	fmt.Println(contenidoArchivo)

}
