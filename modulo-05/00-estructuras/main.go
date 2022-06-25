package main

import "fmt"

//craciond e la clase
type User struct {
	Name  string //atributos
	Email string
}

func (self *User) setName(name string) {
	self.Name = name
}

func (self *User) getName() string { // string porque solo retornara
	return self.Name
}

func main() {
	cody := User{Name: "Cody", Email: "info@mail.com"}

	cody.setName("codigofacilito")

	fmt.Println(cody)
}
