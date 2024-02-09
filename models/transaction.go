package models

type Transaction struct {
	ID             string      `json:"id"`
	Status         string      `json:"status"`
	CompletedAt    *CustomTime `json:"completed_at"`
	PaymentChannel string      `json:"payment_channel"`
}

type TransactionResponse struct {
	BillID       string        `json:"bill_id"`
	Transactions []Transaction `json:"transactions"`
}
