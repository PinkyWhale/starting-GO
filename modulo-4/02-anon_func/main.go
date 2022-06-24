package main

import "fmt"

func main() {

	// Funcion anonima, sin nombre
	miFuncion := func(nombre string) {
		fmt.Println("Hola", nombre, "desde una funcion anonima.")
	}

	miFuncion2 := func(nombresito string) string {
		mensaje := fmt.Sprintf("Hola %s, te saludamos desde una funcion sin nombre", nombresito)
		return mensaje
	}

	miFuncion("Eri")

	msjFinal := miFuncion2("Erica")
	fmt.Println(msjFinal)

}
