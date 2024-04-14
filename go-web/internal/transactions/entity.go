package transactions

type Transaction struct {
	Id       uint64  `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Emiter   string  `json:"emiter"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}
