package main

import (
	"errors"
	"fmt"
)

const minWorkedHours = 80

type salaryError struct {
}

func (s *salaryError) Error() string {
	return "o trabalhador nao pode ter trabalhado menos de 80 horas"
}

func calcMonthlySalary(workedHours float64, hourPrice float64) (float64, error) {
	if workedHours < minWorkedHours {
		return 0, &salaryError{}
	}
	salary := workedHours * hourPrice
	if salary > 20000 {
		afterTaxes := 0.9
		salary = afterTaxes * salary

	}
	return salary, nil
}

func main() {

	_, err := calcMonthlySalary(70.0, 10.0)
	if err != nil {
		err = fmt.Errorf("error: %w", err)
		rawErr := errors.Unwrap(err)
		newErr := errors.New(rawErr.Error())
		fmt.Println(newErr)
	}

	salary, _ := calcMonthlySalary(100.0, 15.0)
	fmt.Println(salary)
}
