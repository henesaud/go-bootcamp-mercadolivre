package main

type Service interface {
	All() ([]Transaction, error)
	Store(id int, code, currency, emiter, receiver, date string, amount float64) (Transaction, error)
}

type service struct {
	respository Repository
}

func NewService(r Repository) Service {
	return &service{respository: r}
}

func (s *service) All() ([]Transaction, error) {
	trn, err := s.respository.All()
	if err != nil {
		return nil, err
	}

	return trn, nil
}

func (s *service) Store(id int, code, currency, emiter, receiver, date string, amount float64) (Transaction, error) {
	lastId, err := s.respository.LastID()
	if err != nil {
		return Transaction{}, err
	}

	lastId++

	trn, err := s.respository.Store(lastId, code, currency, emiter, receiver, date, amount)
	if err != nil {
		return Transaction{}, err
	}

	return trn, nil
}
