package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/internal/transactions"
)

type Request struct {
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Emiter   string  `json:"emiter"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

type TransactionHandler struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *TransactionHandler {
	return &TransactionHandler{
		service: t,
	}
}

func (c *TransactionHandler) All(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	trns, err := c.service.All()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(trns) == 0 {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, trns)
}

func (c *TransactionHandler) Store(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}
	var req Request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(req)
	p, err := c.service.Store(req.Code, req.Currency, req.Emiter, req.Receiver, req.Date, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, p)
}
