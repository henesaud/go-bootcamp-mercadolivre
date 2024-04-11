package main

import (
	"github.com/gin-gonic/gin"
)

var transactions = []Transaction{
	{
		Id:       "1",
		Code:     "99",
		Currency: "USD",
		Amount:   100.0,
		Emiter:   "Hene",
		Receiver: "SS",
		Date:     "2021-01-01",
	},
	{
		Id:       "2",
		Code:     "443",
		Currency: "BRL",
		Amount:   99.0,
		Emiter:   "Caio",
		Receiver: "Souza",
		Date:     "2024-01-01",
	},
}

func All(c *gin.Context) {
	c.JSON(200, transactions)

}

func FindTransactionById(id string) Transaction {
	for _, transaction := range transactions {
		if transaction.Id == id {
			return transaction
		}
	}
	return Transaction{}
}

func TransactionById(c *gin.Context) {
	id := c.Param("id")

	t := FindTransactionById(id)
	if t.Id == "" {
		c.JSON(404, gin.H{"message": "Transaction not found"})
		return
	}

	c.JSON(200, t)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	server := router.Group("/transactions")
	{
		server.GET("/", All)
		server.GET("/:id", TransactionById)
	}

	router.Run()
}
