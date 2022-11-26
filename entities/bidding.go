package entities

type Bidding struct {
	Id        int     `json:"id"`
	BidderId  int     `json:"bidderId"`
	AuctionId int     `json:"auctionId"`
	Amount    float64 `json:"amount"`
	IsActive  bool    `json:"isActive"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
