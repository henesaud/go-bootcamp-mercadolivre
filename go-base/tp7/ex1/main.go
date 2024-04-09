package main

import "fmt"

type user struct {
	name     string
	surname  string
	age      int
	email    string
	password string
}

func (u *user) setName(name string) {
	u.name = name
}

func (u *user) setSurname(surname string) {
	u.surname = surname
}

func (u *user) setAge(age int) {
	u.age = age
}

func (u *user) setEmail(email string) {
	u.email = email
}

func (u *user) setPassword(password string) {
	u.password = password
}

func main() {
	u := user{}
	u.setName("Hene")
	u.setSurname("SS")
	u.setAge(26)
	u.setEmail("hene@ml.com")
	u.setPassword("password123")

	fmt.Println("Name:", u.name)
	fmt.Println("Surname:", u.surname)
	fmt.Println("Age:", u.age)
	fmt.Println("Email:", u.email)
	fmt.Println("Password:", u.password)
}
