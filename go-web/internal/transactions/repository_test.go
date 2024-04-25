package transactions

import (
	"encoding/json"
	"testing"

	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	input := []Transaction{
		{
			Id:       0,
			Code:     "cce223",
			Currency: "BRL",
			Amount:   110.43,
			Emiter:   "Mercado Livre",
			Receiver: "Pedro Costa",
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

	resp, _ := myRepo.All()

	assert.Equal(t, input, resp)
}

func TestUpdate(t *testing.T) {
	id := uint64(1)
	currentTransactions := []Transaction{{
		Id:       id,
		Code:     "oldCode",
		Currency: "oldCurrency",
		Amount:   100.0,
		Emiter:   "oldEmiter",
		Receiver: "oldReceiver",
		Date:     "oldDate",
	},
	}
	newTransaction := Transaction{
		Id:       id,
		Code:     "newCode",
		Currency: "newCurrency",
		Amount:   200.0,
		Emiter:   "newEmiter",
		Receiver: "newReceiver",
		Date:     "newDate",
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

	resp, _ := myRepo.UpdateAmount(id, newTransaction.Amount)
	assert.Equal(t, newTransaction.Amount, resp.Amount)
	assert.Equal(t, true, dbMock.HasCalledRead)
}
