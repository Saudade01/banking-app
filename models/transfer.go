package models

type Transfer struct {
	ID              int64   `json:"id"`
	FromAccountID   int64   `json:"from_account_id"`
	ToAccountID     int64   `json:"to_account_id"`
	Amount          float64 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
}
