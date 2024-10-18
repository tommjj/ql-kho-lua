package domain

import "time"

type InvoiceItem struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	RiceID   int     `json:"rice_id"`
	Rice     *Rice   `json:"rice,omitempty"`
}

type Invoice struct {
	ID           int           `json:"id"`
	StorehouseID int           `json:"storehouse_id"`
	CustomerID   int           `json:"customer_id"`
	UserID       int           `json:"user_id"`
	CreatedAt    time.Time     `json:"created_at"`
	TotalPrice   float64       `json:"total_price"`
	Details      []InvoiceItem `json:"details"`
	CreatedBy    *User         `json:"created_by"`
	Customer     *Customer     `json:"customer"`
	Storehouse   *Storehouse   `json:"storehouse"`
}

// CalcTotalPrice
func (i *Invoice) CalcTotalPrice() float64 {
	for _, v := range i.Details {
		i.TotalPrice += v.Price * float64(v.Quantity)
	}
	return i.TotalPrice
}
