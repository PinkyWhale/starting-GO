package main

import "fmt"

// Variadic Function --- por que podemos poner la n cantidad de argumentos
func promedio(calificaciones ...int) int {

	var promedio, suma int

	for _, calicalificacion := range calificaciones {
		suma = suma + calicalificacion
	}

	promedio = suma / len(calificaciones)

	return promedio

}
func main() {
	resultado := promedio(10, 10, 10, 9, 10)
	fmt.Println("Promedio:", resultado)
}
