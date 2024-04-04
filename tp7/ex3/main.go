package main

import (
	"fmt"
)

type product struct {
	name     string
	price    float64
	quantity int
}

type services struct {
	name          string
	price         float64
	workedMinutes int
}

type maintenance struct {
	name  string
	price float64
}

func sumProducts(products []product) float64 {
	var sum float64
	for _, p := range products {
		sum += p.price * float64(p.quantity)
	}
	return sum
}

func sumServices(s []services) float64 {
	var sum float64

	var workedHours float64
	for _, s := range s {
		workedHours = 0.5
		if s.workedMinutes > 30 {
			workedHours = float64(s.workedMinutes) / 60.0
		}
		sum += s.price * float64(workedHours)
	}
	return sum
}

func sumMaintenances(m []maintenance) float64 {
	var sum float64
	for _, m := range m {
		sum += m.price
	}
	return sum
}

func main() {
	products := []product{
		{name: "Product 1", price: 10.0, quantity: 2},
		{name: "Product 2", price: 5.0, quantity: 3},
		{name: "Product 3", price: 8.0, quantity: 1},
	}

	services := []services{
		{name: "Service 1", price: 20.0, workedMinutes: 125},
		{name: "Service 2", price: 15.0, workedMinutes: 70},
		{name: "Service 3", price: 12.0, workedMinutes: 31},
		{name: "Service 3", price: 100.0, workedMinutes: 11},
	}

	maintenances := []maintenance{
		{name: "Maintenance 1", price: 50.0},
		{name: "Maintenance 2", price: 30.0},
		{name: "Maintenance 3", price: 40.0},
	}

	sumProductsCh := make(chan float64)
	sumServicesCh := make(chan float64)
	sumMaintenancesCh := make(chan float64)

	go func() {
		sum := sumProducts(products)
		sumProductsCh <- sum
	}()

	go func() {
		sum := sumServices(services)
		sumServicesCh <- sum
	}()

	go func() {
		sum := sumMaintenances(maintenances)
		sumMaintenancesCh <- sum
	}()

	total := <-sumProductsCh + <-sumServicesCh + <-sumMaintenancesCh
	fmt.Println("Total:", total)
}
