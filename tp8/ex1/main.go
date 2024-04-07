package main

import (
	"fmt"
	"os"
)

type MyError struct {
	minSalary int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("salary must be greater than %d", e.minSalary)
}

func myCustomErrorTest(salary int) error {
	if salary < 15000 {
		return &MyError{minSalary: salary}
	}
	return nil
}

func main() {
	var salary int = 10000
	ok := myCustomErrorTest(salary)
	if ok != nil {
		fmt.Println("error:", ok)
		os.Exit(1)
	}
	fmt.Println("You must pay tax!")
}
