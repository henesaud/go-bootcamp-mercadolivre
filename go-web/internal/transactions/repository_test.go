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
