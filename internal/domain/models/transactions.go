package models

import "time"

type Transaction struct {
	Id            string    `json:"id"`
	UserID        string    `json:"user_id"`
	CategoryID    string    `json:"category_id"`
	Amount        int       `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Date          time.Time `json:"date"`
	Note          string    `json:"note"`
	Created_at    time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
