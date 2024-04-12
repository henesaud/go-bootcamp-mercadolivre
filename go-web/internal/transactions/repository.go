package main

type Transaction struct {
	Id       int     `json:"name"`
	Code     string  `json:"age"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Emiter   string  `json:"emiter"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

var trn []Transaction
var lastId int

type Repository interface {
	All() ([]Transaction, error)
	Store(id int, code, currency, emiter, receiver, date string, amount float64) (Transaction, error)
	LastID() (int, error)
}

type repository struct{}

func (r *repository) All() ([]Transaction, error) {
	return trn, nil
}

func (r *repository) Store(id int, code, currency, emiter, receiver, date string, amount float64) (Transaction, error) {
	p := Transaction{
		Id:       id,
		Code:     code,
		Currency: currency,
		Amount:   amount,
		Emiter:   emiter,
		Receiver: receiver,
		Date:     date,
	}
	trn = append(trn, p)
	lastId = p.Id
	return p, nil
}

func (r *repository) LastID() (int, error) {
	return lastId, nil
}

func NewRepository() Repository {
	return &repository{}
}
