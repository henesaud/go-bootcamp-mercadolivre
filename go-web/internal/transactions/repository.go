package transactions

type Repository interface {
	All() ([]Transaction, error)
	Store(code, currency, emiter, receiver, date string, amount float64) (Transaction, error)
	LastID() (uint64, error)
	Update(id uint64, code, currency, emiter, receiver, date string, amount float64) (Transaction, error)
	UpdateAmount(id uint64, amount float64) (Transaction, error)
	Delete(id uint64) error
}

func NewRepository() Repository {
	return &MemoryRepository{}
}
