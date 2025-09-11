package models

import "time"

type User struct {
	Id         string    `json:"id"`
	Phone      string    `json:"phone_number"`
	Name       string    `json:"name"`
	Currency   string    `json:"currency"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
