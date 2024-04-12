package main

import (
	"github.com/gin-gonic/gin"
)

// var transactions = []Transaction{
// 	{
// 		Id:       "1",
// 		Code:     "99",
// 		Currency: "USD",
// 		Amount:   100.0,
// 		Emiter:   "Hene",
// 		Receiver: "SS",
// 		Date:     "2021-01-01",
// 	},
// 	{
// 		Id:       "2",
// 		Code:     "443",
// 		Currency: "BRL",
// 		Amount:   99.0,
// 		Emiter:   "Caio",
// 		Receiver: "Souza",
// 		Date:     "2024-01-01",
// 	},
// }

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run()
}
