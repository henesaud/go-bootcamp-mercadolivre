// Aula 01

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Transaction struct {
	Id       string  `json:"name"`
	Code     string  `json:"age"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Emiter   string  `json:"emiter"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

func CreateTransactionJson() {
	t := Transaction{
		Id:       "1",
		Code:     "99",
		Currency: "USD",
		Amount:   100.0,
		Emiter:   "Hene",
		Receiver: "SS",
		Date:     "2021-01-01",
	}

	tJson, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("transactions.json", tJson, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JSON file created successfully.")
}
