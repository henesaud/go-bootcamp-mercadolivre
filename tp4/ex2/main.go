package main

import "fmt"

func calcMean(values ...float64) (float64, error) {
	sum := 0.0
	for _, n := range values {
		if n < 0.0 {
			return 0.0, fmt.Errorf("negative values are not allowed: %.2f", n)
		}
		sum += n
	}
	return sum / float64(len(values)), nil
}

func main() {
	mean, err := calcMean(3, 4.5, 3, 9, 10, 8.5)
	if err != nil {
		fmt.Println("Error.", err)
		return
	}
	fmt.Printf("Mean: %.2f\n", mean)
}
