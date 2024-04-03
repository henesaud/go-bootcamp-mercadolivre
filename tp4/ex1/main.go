package main

import "fmt"

func calculateTaxe(salary float64) float64 {
	if salary <= 50000.0 {
		return 0.0
	} else if salary <= 150000.0 {
		return salary * 0.17
	}
	return salary * 0.27
}

func main() {
	fmt.Printf("imposto: até R$50.000: %.2f\n", calculateTaxe(50000))
	fmt.Printf("imposto: até R$150.000: %.2f\n", calculateTaxe(150000))
	fmt.Printf("imposto: maior que R$150.000: %.2f\n", calculateTaxe(150001))
}
