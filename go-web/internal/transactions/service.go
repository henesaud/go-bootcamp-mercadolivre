package transactions

type Service interface {
	All() ([]Transaction, error)
	Store(code, currency, emiter, receiver, date string, amount float64) (Transaction, error)
}

type service struct {
	repository Repository
}

func (s *service) All() ([]Transaction, error) {
	trans, err := s.repository.All()
	if err != nil {
		return nil, err
	}

	return trans, nil

}

func (s *service) Store(code, currency, emiter, receiver, date string, amount float64) (Transaction, error) {
	trans, err := s.repository.Store(code, currency, emiter, receiver, date, amount)
	if err != nil {
		return Transaction{}, err
	}

	return trans, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
