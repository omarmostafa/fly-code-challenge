package Entities

type TransactionEntity struct {
	Id             string `json:"id"`
	Amount         int    `json:"amount"`
	StatusCode     string `json:"status_code"`
	OrderReference string `json:"order_reference"`
	Currency       string `json:"currency"`
}
