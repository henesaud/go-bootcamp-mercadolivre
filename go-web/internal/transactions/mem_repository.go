package transactions

import "fmt"

var trans []Transaction
var lastID uint64 = 0

type MemoryRepository struct {
}

func (m *MemoryRepository) All() ([]Transaction, error) {
	return trans, nil
}

func (m *MemoryRepository) Store(code, currency, emiter, receiver, date string, amount float64) (Transaction, error) {
	lastID++
	p := Transaction{
		Id:       lastID,
		Code:     code,
		Currency: currency,
		Amount:   amount,
		Emiter:   emiter,
		Receiver: receiver,
		Date:     date,
	}
	trans = append(trans, p)
	return p, nil
}

func (m *MemoryRepository) LastID() (uint64, error) {
	return lastID, nil
}

func (m *MemoryRepository) Update(id uint64, code, currency, emiter, receiver, date string, amount float64) (Transaction, error) {
	t := Transaction{
		Code:     code,
		Currency: currency,
		Amount:   amount,
		Emiter:   emiter,
		Receiver: receiver,
		Date:     date,
	}
	updated := false
	for i := range trans {
		if trans[i].Id == id {
			t.Id = id
			trans[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaction{}, transactionNotFoundError(id)
	}
	return t, nil
}

func (m *MemoryRepository) UpdateAmount(id uint64, amount float64) (Transaction, error) {
	var t Transaction
	updated := false
	for i := range trans {
		if trans[i].Id == id {
			trans[i].Amount = amount
			updated = true
			t = trans[i]
		}
	}
	if !updated {
		return Transaction{}, transactionNotFoundError(id)
	}

	return t, nil
}

func (r *MemoryRepository) Delete(id uint64) error {
	deleted := false
	var index int
	for i := range trans {
		if trans[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return transactionNotFoundError(id)
	}

	trans = append(trans[:index], trans[index+1:]...)
	return nil
}

func transactionNotFoundError(id uint64) error {
	return fmt.Errorf("transaction %d not found", id)
}
