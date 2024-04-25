package transactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

var newTransaction = Transaction{
	Id:       1,
	Code:     "newCode",
	Currency: "newCurrency",
	Amount:   200.0,
	Emiter:   "newEmiter",
	Receiver: "newReceiver",
	Date:     "newDate",
}

func TestServiceAll(t *testing.T) {
	input := []Transaction{
		{
			Id:       1,
			Code:     "cce223",
			Currency: "BRL",
			Amount:   110.43,
			Emiter:   "Empresa1",
			Receiver: "Caio dias",
		},
		{
			Id:       2,
			Code:     "cce5534",
			Currency: "USD",
			Amount:   143.6,
			Emiter:   "Mercado Pago",
			Receiver: "Abdullah",
		},
	}

	dataJson, _ := json.Marshal(input)

	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.All()

	assert.Equal(t, input, result)
	assert.Nil(t, err)
}

func TestServiceAllError(t *testing.T) {
	expectedError := errors.New("error for All")

	dbMock := store.Mock{
		Err: expectedError,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.All()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)

}

func TestServiceUpdate(t *testing.T) {
	currentTransactions := []Transaction{
		{
			Id:       1,
			Code:     "cce223",
			Currency: "BRL",
			Amount:   110.43,
			Emiter:   "Empresa1",
			Receiver: "Caio dias",
		},
		{
			Id:       2,
			Code:     "cce5534",
			Currency: "USD",
			Amount:   143.6,
			Emiter:   "Mercado Pago",
			Receiver: "Abdullah",
		},
	}

	dataJson, _ := json.Marshal(currentTransactions)

	dbMock := store.Mock{
		Data:          dataJson,
		HasCalledRead: false,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.Update(newTransaction.Id, newTransaction.Code, newTransaction.Currency, newTransaction.Emiter, newTransaction.Receiver, newTransaction.Date, newTransaction.Amount)

	assert.Equal(t, newTransaction, result)
	assert.Nil(t, err)
	assert.Equal(t, true, dbMock.HasCalledRead)
}

func TestDelete(t *testing.T) {
	currentTransactions := []Transaction{
		{
			Id:       4,
			Code:     "cce223",
			Currency: "BRL",
			Amount:   110.43,
			Emiter:   "Empresa1",
			Receiver: "Caio dias",
		},
	}

	dataJson, _ := json.Marshal(currentTransactions)

	dbMock := store.Mock{
		Data:          dataJson,
		HasCalledRead: false,
	}

	storeStub := store.FileStoreMock{
		FileName: "mockedFile",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	err := myService.Delete(currentTransactions[0].Id)
	assert.Nil(t, err)

	nonExistingId := 38
	err = myService.Delete(uint64(nonExistingId))
	expectedErrorMessage := fmt.Sprintf("transaction %d not found", nonExistingId)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErrorMessage, err.Error())

}

func TestStore(t *testing.T) {
	testTransaction := Transaction{
		Id:       1,
		Code:     "cce223",
		Currency: "BRL",
		Amount:   110.43,
		Emiter:   "Empresa1",
		Receiver: "Caio dias",
	}

	dbMock := store.Mock{
		Data: []byte("[]"),
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.Store(testTransaction.Code, testTransaction.Currency, testTransaction.Emiter, testTransaction.Receiver, testTransaction.Date, testTransaction.Amount)

	assert.Equal(t, testTransaction.Code, result.Code)
	assert.Equal(t, testTransaction.Currency, result.Currency)
	assert.Equal(t, testTransaction.Emiter, result.Emiter)
	assert.Equal(t, testTransaction.Receiver, result.Receiver)
	assert.Equal(t, testTransaction.Date, result.Date)
	assert.Equal(t, testTransaction.Amount, result.Amount)

}

func TestStoreError(t *testing.T) {
	testProduct := Transaction{
		Id:       2,
		Code:     "cce5534",
		Currency: "USD",
		Amount:   143.6,
		Emiter:   "Mercado Pago",
		Receiver: "Abdullah",
	}
	expectedError := errors.New("error for Storage")

	dbMock := store.Mock{
		Err: expectedError,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.Store(testProduct.Code, testProduct.Currency, testProduct.Emiter, testProduct.Receiver, testProduct.Date, testProduct.Amount)

	assert.Equal(t, Transaction{}, result)
	assert.Equal(t, expectedError, err)

}
