package transactions

type Repository interface {
	All() ([]Transaction, error)
	Store(code, currency, emiter, receiver, date string, amount float64) (Transaction, error)
	LastID() (uint64, error)
}

func NewRepository() Repository {
	return &MemoryRepository{}
}
