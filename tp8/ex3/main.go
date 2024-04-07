package main

import (
	"fmt"
	"os"
)

const minSalary = 15000

func main() {
	var salary int = 10000
	if salary < minSalary {
		fmt.Println(fmt.Errorf("erro: o minimo tributavel e 15.000 e o salario informado e: %d", salary))
		os.Exit(1)
	}
	fmt.Println("You must pay tax!")
}
