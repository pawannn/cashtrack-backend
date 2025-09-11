package models

import "time"

type Category struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Icon      string    `json:"icon"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
