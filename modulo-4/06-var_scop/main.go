package main

import "fmt"

func modificarVariable(parametro string) {
	fmt.Println("Valor actual:", parametro)
	parametro = "Nuevo Curso de GO!"
	fmt.Println("Nuevo valor:", parametro)

	fmt.Printf("%p\n", &parametro)
}

func main() {

	var curso = "Curso Profesional de go!"

	modificarVariable(curso)

	fmt.Println(">>>", curso)
	fmt.Printf("%p\n", &curso)

}
