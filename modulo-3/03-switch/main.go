package main

import "fmt"

func main() {

	var calificacion int
	fmt.Println("Ingresa una calificacion: ")
	fmt.Scanf("%d", &calificacion)

	switch calificacion {
	case 10:
		fmt.Println("Felicidades, obtuviste una calificacion Perfecta!")
	case 8, 9:
		fmt.Println("Aprobaste la Materia.")
	case 7, 6:
		fmt.Println("Aprobaste la materia, pero necesitas estudiar más.")
	case 1, 2, 3, 4, 5:
		fmt.Println("No aprobaste la Materia.")
	default:
		fmt.Println("Calificación no valida.")
	}
}
