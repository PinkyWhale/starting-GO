package main

import "fmt"

func main() {

	// break - continue

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue // salta iteracion
		}

		if i == 8 {
			break // finaliza la ejecucion del ciclo
		}

		fmt.Println(i)
	}
}
