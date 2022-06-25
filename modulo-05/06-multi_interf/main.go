package main

import "fmt"

type Animal interface {
	Domir()
}

type Mascota interface {
	Comer()
}

type Gato struct {
	Nombre string
}

func (self Gato) Domir() {
	fmt.Println("El gato duerme")
}
func (self Gato) Comer() {
	fmt.Println("El gato duerme")
}

func funcionParaUnAnimal(animal Animal) {
	fmt.Println("El objeto es un animal")
}

func funcionParaUnaMascota(mascota Mascota) {
	fmt.Println("El objeto es un animal")
}

func main() {
	gato := Gato{Nombre: "Mi gato!"}

	gato.Domir()
	gato.Comer()
}
