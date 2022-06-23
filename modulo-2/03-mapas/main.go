package main

import "fmt"

func main() {
	// Los maps son arreglos asociados/ diccionarios en python
	dias := make(map[int]string)

	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"

	delete(dias, 4) // se borra la llave 4 con su correspondiente valor

	fmt.Println(dias)

	usuarios := make(map[string][]int)

	usuarios["Eduardo"] = []int{10, 9, 8, 10, 5}

	fmt.Println(usuarios)
}
