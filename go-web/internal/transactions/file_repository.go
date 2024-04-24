package transactions

import (
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/pkg/store"
)

type FileRepository struct {
	db store.Store
}

func NewFileRepository(db store.Store) Repository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) All() ([]Transaction, error) {
	var ps []Transaction
	err := r.db.Read(&ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (r *FileRepository) Store(code, currency, emiter, receiver, date string, amount float64) (Transaction, error) {
	t := Transaction{
		Code:     code,
		Currency: currency,
		Emiter:   emiter,
		Receiver: receiver,
		Date:     date,
		Amount:   amount,
	}

	var trns []Transaction

	r.db.Read(&trns)

	lastIdInserted := len(trns)
	lastIdInserted++
	t.Id = uint64(lastIdInserted)

	trns = append(trns, t)

	err := r.db.Write(trns)
	if err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func (r *FileRepository) Delete(id uint64) error {
	var trns []Transaction

	r.db.Read(&trns)

	deleted := false
	var index int
	for i := range trns {
		if trns[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return transactionNotFoundError(id)
	}
	trns = append(trns[:index], trns[index+1:]...)

	err := r.db.Write(trns)
	if err != nil {
		return err
	}
	return nil
}

func (r *FileRepository) Update(id uint64, code, currency, emiter, receiver, date string, amount float64) (Transaction, error) {
	var trns []Transaction
	r.db.Read(&trns)

	t := Transaction{
		Code:     code,
		Currency: currency,
		Amount:   amount,
		Emiter:   emiter,
		Receiver: receiver,
		Date:     date,
	}

	updated := false
	for i := range trns {
		if trns[i].Id == id {
			t.Id = id
			trns[i] = t
			updated = true
		}
	}
	if !updated {
		return Transaction{}, transactionNotFoundError(id)
	}

	err := r.db.Write(trns)
	if err != nil {
		return Transaction{}, err
	}
	return t, nil
}
func (r *FileRepository) UpdateAmount(id uint64, amount float64) (Transaction, error) {
	var trns []Transaction
	r.db.Read(&trns)

	var t Transaction
	updated := false
	for i := range trns {
		if trns[i].Id == id {
			trns[i].Amount = amount
			updated = true
			t = trns[i]
		}
	}
	if !updated {
		return Transaction{}, transactionNotFoundError(id)
	}

	err := r.db.Write(trns)
	if err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func (r *FileRepository) LastID() (uint64, error) {
	var ps []Transaction
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil

}
