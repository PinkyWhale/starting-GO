package main

import "fmt"

func main() {

	var calificacion int
	fmt.Println("Ingresa una calificacion: ")
	fmt.Scanf("%d", &calificacion)

	if calificacion == 10 {
		fmt.Println("Felicidades, obtuviste una calificacion Perfecta!")
	} else if calificacion == 8 || calificacion == 9 {
		fmt.Println("Aprobaste la Materia.")
	} else if calificacion == 6 || calificacion == 7 {
		fmt.Println("Aprobaste la materia, pero necesitas estudiar más.")
	} else if calificacion >= 0 && calificacion <= 5 {
		fmt.Println("No aprobaste la Materia.")
	} else {
		fmt.Println("Calificación no valida.")
	}
}
