package entities

import "time"

type Auction struct {
	Id        int       `json:"id"`
	AdSpaceId int       `json:"adSpaceId"`
	EndTime   time.Time `json:"endTime"`
	Status    string    `json:"status"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
