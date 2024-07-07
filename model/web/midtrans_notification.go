package web

type MidtransNotification struct {
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionID     string `json:"transaction_id"`
	StatusCode        string `json:"status_code"`
	GrossAmount       string `json:"gross_amount"`
	Currency          string `json:"currency"`
}
