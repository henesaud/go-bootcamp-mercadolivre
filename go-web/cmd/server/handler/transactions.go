package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

type PartialRequest struct {
	Amount float64 `json:"amount"`
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

func (c *TransactionHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(req.Amount, req.Code, req.Currency, req.Date, req.Emiter, req.Receiver)

	if req.Code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "code is a must"})
		return
	}

	if req.Currency == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "curency is a must"})
		return
	}

	if req.Amount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "amount is a must"})
		return
	}

	if req.Emiter == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "emiter is a must"})
		return
	}

	if req.Receiver == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "receiver is a must"})
		return
	}

	if req.Date == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "date is a must"})
		return
	}

	p, err := c.service.Update(id, req.Code, req.Currency, req.Emiter, req.Receiver, req.Date, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, p)
}

func (c *TransactionHandler) UpdateAmount(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	var req PartialRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if req.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "the amount must be greater than 0"})
		return
	}
	t, err := c.service.UpdateAmount(id, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, t)
}

func (c *TransactionHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("transaction %d removed", id)})
}
