package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Customer struct {
	fileId,
	name,
	surname,
	rg,
	cellphone,
	address string
}

func genId() (string, error) {
	return strconv.Itoa(rand.Intn(1000000)), nil
}

func treatOpenFile(fileName string) *os.File {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	return OpenFile(fileName)
}

func genNewCustomer() *Customer {
	id, err := genId()
	if err != nil {
		panic("ocorreu um erro ao tentar gerar id")
	}

	newClient := Customer{
		fileId:    id,
		name:      "HeneNovo",
		surname:   "SS",
		rg:        "501548889",
		cellphone: "31990707060",
		address:   "endereco qualquer",
	}

	return &newClient
}

func findCustomer(id string, customers []Customer) *Customer {
	for i := 0; i < len(customers); i++ {
		if customers[i].fileId == id {
			return &customers[i]
		}
	}
	return nil
}

func main() {
	file := treatOpenFile("customers.txt")
	err := errors.New("")
	if file == nil {
		file, err = os.Create("customers.txt")
		if err != nil {
			panic("nao foi possivel criar o arquivo")
		}
	}
	scanner := bufio.NewScanner(file)
	readData := make([]Customer, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		readData = append(readData, Customer{
			fileId:    parts[0],
			name:      parts[1],
			surname:   parts[2],
			rg:        parts[3],
			cellphone: parts[4],
			address:   parts[5],
		})
	}

	newCustomer := genNewCustomer()
	newCustomerExists := findCustomer(newCustomer.fileId, readData)

	var newData []Customer
	if newCustomerExists == nil {
		newData = append(readData, *newCustomer)
	}
	fmt.Println(newData)

	for _, customer := range newData {
		currentLine := fmt.Sprintf("\n%s,%s,%s,%s,%s,%s", customer.fileId, customer.name, customer.surname, customer.rg, customer.cellphone, customer.address)
		_, err := file.WriteString(currentLine)
		if err != nil {
			fmt.Println(err)
		}
	}

}
