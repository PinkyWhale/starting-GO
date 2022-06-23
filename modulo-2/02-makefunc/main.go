package main

import "fmt"

func main() {
	slice := make([]int, 3, 3)
	/*
		1) el arreglo
		2) la longitud
		3) capacidad maxima
	*/

	slice[0] = 100
	slice[1] = 200
	slice[2] = 300
	slice[3] = 400

	slice = append(slice, 400)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
