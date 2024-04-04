package main

import (
	"bufio"
	"fmt"
	"os"
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
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for _, v := range values {
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}
