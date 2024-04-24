package transactions

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

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

	assert.Equal(t, expectedError, err)
	assert.Equal(t, Transaction{}, result)

}
