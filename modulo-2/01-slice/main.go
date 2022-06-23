package main

import "fmt"

func main() {

	numeros := []int{1, 2, 3, 4}

	// incrementacion
	numeros = append(numeros, 5)

	fmt.Println(numeros)

	//nuevo slice
	nuevoSlice := numeros[0:5] //un slice es una referencia de un arreglo ya definido
	numeros[0] = 100

	fmt.Println(nuevoSlice)

	//----- Practica

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre"}

	meses = append(meses, "Octubre") // Si la estructura se encuentra en su capacidad maxima se genera un nuevo arreglo.

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Println("La longitud es: %v. La capacidad es: %v %p\n", longitud, capacidad, meses)

	//RESUMEN
	nums := []int{1, 2, 3, 4, 5, 6}

	//primer slice
	inicio := nums[0:3]
	//segundo slice
	final := nums[3:6]

	fmt.Println(nums)
	fmt.Println(inicio)
	fmt.Println(final)

}
