package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/cmd/server/handler"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/internal/transactions"
)

func main() {
	rep := transactions.NewRepository()
	service := transactions.NewService(rep)
	trnsHandler := handler.NewTransaction(service)

	server := gin.Default()
	trns := server.Group("/transactions")
	trns.POST("/", trnsHandler.Store)
	trns.GET("/", trnsHandler.All)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
