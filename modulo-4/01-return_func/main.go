package main

import "fmt"

//parametros que recibe y tipo + parametros que devuelve y tipo
func suma(num1, num2 int) (int, string) {

	return num1 + num2, "Un mensaje desde la consola."

}

func main() {

	resultado, mensaje := suma(10, 20) // si no quiero recibir alguno lo reemplazo por __
	fmt.Println(resultado, mensaje)
}
