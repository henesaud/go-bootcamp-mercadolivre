package transactions

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
