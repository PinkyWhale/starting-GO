package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

func main() {

	eduardo := User{Name: "Eduardo", Email: "unmail@mail.com", Active: true}

	uriel := User{Name: "Uriel", Email: "unmail2@mail.com", Active: true}

	eduardoStudent := Student{User: eduardo, Id: "1f1"}

	fmt.Println(uriel)
	fmt.Println(eduardoStudent.User.Name)
}
