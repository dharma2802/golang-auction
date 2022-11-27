package entities

import "time"

type AdSpace struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Position  string    `json:"position"`
	Width     float64   `json:"width"`
	Height    float64   `json:"height"`
	Price     float64   `json:"price"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
