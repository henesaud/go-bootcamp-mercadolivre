package main

import (
	"errors"
	"fmt"
	"os"
)

const minSalary = 15000

func main() {
	var salary int = 10000
	if salary < minSalary {
		errorMessage := fmt.Sprintf("salary must be greater than %d", minSalary)
		fmt.Println(errors.New(errorMessage))
		os.Exit(1)
	}
	fmt.Println("You must pay tax!")
}
