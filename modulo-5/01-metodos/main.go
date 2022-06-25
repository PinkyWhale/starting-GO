package main

import "fmt"

//craciond e la clase
type User struct {
	Name  string //atributos
	Email string
}

func main() {

	Name := "Cody"
	Email := "unmail@mail.com"

	//creo el objeto
	cody := User{Name, Email}

	fmt.Println(cody.Name)
	fmt.Println(cody.Email)
}
