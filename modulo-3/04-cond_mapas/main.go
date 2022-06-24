package main

import "fmt"

func main() {
	usuarios := make(map[string]string)

	usuarios["Eduardo"] = "eduardo@codigofacilito.com"

	correo, ok := usuarios["Eduardo"]
	// retorna 2 valores, el original y la segunda opcion es un booleando... ok == true

	if ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
