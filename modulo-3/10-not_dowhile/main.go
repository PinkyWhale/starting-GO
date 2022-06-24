package main

import "fmt"

func main() {

	//DO While

	var numero = 10

	//--- inicializamos en true; condicionamos (mientras que verdadero); indicamos sentencia (cuando queremos que finalice)
	for ok := true; ok; ok = numero < 10 {
		fmt.Println("numero")
		numero++
	}

}
