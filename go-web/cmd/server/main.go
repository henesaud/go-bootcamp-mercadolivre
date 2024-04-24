package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/cmd/server/handler"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/internal/transactions"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/pkg/store"
	"github.com/joho/godotenv"
)

func LoggerMiddleware(ctx *gin.Context) {
	fmt.Printf("[%s] %s\n", ctx.Request.Method, ctx.Request.URL)
	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	tokenEnvironment := os.Getenv("TOKEN")
	token := ctx.GetHeader("token")
	if token != tokenEnvironment {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}
	ctx.Next()
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	db := store.NewFileStore("file", "transactions.json")

	rep := transactions.NewRepository(db)
	service := transactions.NewService(rep)
	trnsHandler := handler.NewTransaction(service)

	server := gin.Default()
	server.Use(LoggerMiddleware)

	trns := server.Group("/transactions")
	trns.Use(TokenMiddleware)
	trns.POST("/", trnsHandler.Store)
	trns.GET("/", trnsHandler.All)
	trns.PUT("/:id", trnsHandler.Update)
	trns.PATCH("/:id", trnsHandler.UpdateAmount)
	trns.DELETE("/:id", trnsHandler.Delete)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
