package entities

import "time"

type Bidding struct {
	Id        int       `json:"id"`
	BidderId  int       `json:"bidderId"`
	AuctionId int       `json:"auctionId"`
	Amount    float64   `json:"amount"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
