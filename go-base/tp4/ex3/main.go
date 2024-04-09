package main

import (
	"errors"
	"fmt"
)

func calculateSalary(category string, hours int) (float64, error) {
	if category == "C" {
		return float64(hours) * 1000, nil
	}

	if category == "B" {
		salary := float64(hours) * 1500
		if hours > 160 {
			return salary * 1.2, nil
		}
		return salary, nil
	}

	if category == "A" {
		salary := float64(hours) * 3000
		if hours > 160 {
			return salary * 1.5, nil
		}
		return salary, nil
	}

	return 0.0, errors.New("invalid category")

}

func main() {
	salary, _ := calculateSalary("C", 162)
	fmt.Printf("Salary 1: %.2f\n", salary)
	salary, _ = calculateSalary("B", 176)
	fmt.Printf("Salary 2: %.2f\n", salary)
	salary, _ = calculateSalary("A", 172)
	fmt.Printf("Salary 3: %.2f\n", salary)
}
