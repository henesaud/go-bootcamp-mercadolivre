package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func dogFoodQuantidy(quantity int) int {
	return quantity * 10000
}

func catFoodQuantidy(quantity int) int {
	return quantity * 5000
}

func hamsterFoodQuantidy(quantity int) int {
	return quantity * 250
}

func tarantulaFoodQuantidy(quantity int) int {
	return quantity * 150
}

func Animal(t string) (func(quantity int) int, error) {
	if t == dog {
		return dogFoodQuantidy, nil
	}
	if t == cat {
		return catFoodQuantidy, nil
	}
	if t == hamster {
		return hamsterFoodQuantidy, nil
	}
	if t == tarantula {
		return tarantulaFoodQuantidy, nil
	}
	return nil, errors.New("invalid animal type")
}

func main() {
	dfunc, _ := Animal(dog)
	fmt.Printf("cachorro: %d gramas\n", dfunc(3))
	cfunc, _ := Animal(cat)
	fmt.Printf("gato: %d gramas\n", cfunc(1))
	tfunc, _ := Animal(tarantula)
	fmt.Printf("tarantula: %d gramas\n", tfunc(4))
	_, err := Animal("invalid animal")
	if err != nil {
		fmt.Println("Error.", err)
	}
}
