package models

type Account struct {
	ID       int64   `json:"id"`
	Owner    string  `json:"owner"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
