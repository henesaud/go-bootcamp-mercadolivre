package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "min"
	average = "avg"
	maximum = "max"
)

func findMin(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("no input values")
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min, nil
}

func findMax(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("no input values")
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max, nil
}

func calcAvg(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("no input values")
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values)), nil
}

func getMathFunc(t string) (func(values ...float64) (float64, error), error) {
	if t == minimum {
		return findMin, nil
	}
	if t == average {
		return calcAvg, nil
	}
	if t == maximum {
		return findMax, nil
	}
	return nil, errors.New("invalid function type")
}

func main() {
	avgFunc, _ := getMathFunc(average)
	maxFunc, _ := getMathFunc(maximum)
	minFunc, _ := getMathFunc(minimum)
	values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min, _ := minFunc(values...)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(values...)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(values...)
	fmt.Printf("max: %.2f\n", max)
}
