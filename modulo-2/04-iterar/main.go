package main

import "fmt"

func main() {

	usuarios := map[int]string{} // esto es como el make pero en iteracion

	usuarios[1] = "Usuario 1"
	usuarios[2] = "Usuario 2"
	usuarios[3] = "Usuario 3"
	usuarios[4] = "Usuario 4"

	//ciclo for each
	for llave, valor := range usuarios {
		fmt.Println(llave, valor)
	}
}
