package main

import "fmt"

func exercise1() {
	name := "hene"
	age := 26
	fmt.Println(name, age)
}
func exercise2() {
	var temperature float32 = 25.0
	var moisture float32 = 0.35
	var pressure float32 = 1019
	fmt.Println(temperature, moisture, pressure)
}

func exercise3() {
	var nome string
	var sobrenome string
	var idade int
	sobrenome = "Silva"
	var licensa_para_dirigir = true
	var statura_da_pessoa int
	quantidadeDeFilhos := 2
	fmt.Println(nome, sobrenome, idade, licensa_para_dirigir, statura_da_pessoa, quantidadeDeFilhos)
}

func exercise4() {
	var sobrenome string = "Silva"
	var idade int = 25
	boolean := false
	var salario float32 = 4585.90
	var nome string = "Fellipe"
	fmt.Println(sobrenome, idade, boolean, salario, nome)
}
func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}
