package main

import (
	"fmt"
	"os"
)

type product struct {
	id       int
	quantity int
	price    float64
}

func (p product) CSVLine() string {
	return fmt.Sprintf("%d,%d,%.2f\n", p.id, p.quantity, p.price)
}

func (p product) CSVHeader() string {
	return "id,quant,price\n"
}

func genCSV(path string, products []product) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("error when oppening file: %w", err)
	}

	defer file.Close()
	p := products[0]

	if _, err = file.WriteString(p.CSVHeader()); err != nil {
		return fmt.Errorf("error when generating header: %w", err)
	}

	for _, p = range products {
		if _, err = file.WriteString(p.CSVLine()); err != nil {
			return fmt.Errorf("error when saving line: %w", err)
		}
	}

	return nil
}

func main() {
	p := []product{
		{
			id:       0,
			quantity: 10,
			price:    13.9,
		},

		{
			id:       2,
			quantity: 20,
			price:    70.9,
		},
		{
			id:       5,
			quantity: 30,
			price:    991.4,
		},
	}

	genCSV("products.csv", p)

}
