package domain

type InvoiceItem struct {
	Rice     Rice    `json:"rice"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Invoice struct {
	ID         int           `json:"id"`
	UserID     int           `json:"user_id"`
	CustomerID int           `json:"customer_id"`
	Details    []InvoiceItem `json:"details"`
	CreatedBy  User          `json:"created_by"`
	Customer   Customer      `json:"customer"`
}
