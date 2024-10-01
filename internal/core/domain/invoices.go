package domain

type InvoiceItem struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	RiceID   int     `json:"rice_id"`
	Rice     *Rice   `json:"rice,omitempty"`
}

type Invoice struct {
	ID           int           `json:"id"`
	UserID       int           `json:"user_id"`
	CustomerID   int           `json:"customer_id"`
	StorehouseID int           `json:"storehouse_id"`
	TotalPrice   float64       `json:"total_price"`
	Details      []InvoiceItem `json:"details"`
	CreatedBy    *User         `json:"created_by"`
	Customer     *Customer     `json:"customer"`
	Storehouse   *Storehouse   `json:"storehouse"`
}
