package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main() {
	file, err := os.Open("../ex1/products.csv")
	if err != nil {
		fmt.Println(fmt.Errorf("error opening file"))
	}
	defer file.Close()

	w := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t', tabwriter.AlignRight)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	header := strings.Split(scanner.Text(), ",")
	for _, c := range header {
		fmt.Fprintf(w, "%s\t", c)
	}
	fmt.Fprintln(w)
	sum := 0.0
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for idx, v := range values {
			quantity, err := strconv.ParseFloat(values[1], 64)
			if err != nil {
				fmt.Println(fmt.Errorf("error parsing float"))
			}

			if idx == 2 {
				price, err := strconv.ParseFloat(v, 64)
				if err != nil {
					fmt.Println(fmt.Errorf("error parsing float"))
				}
				sum += price * float64(quantity)
			}
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintf(w, "\t\t%.2f \n", sum)

	w.Flush()
}
