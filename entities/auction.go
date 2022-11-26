package entities

type Auction struct {
	Id        int    `json:"id"`
	AdSpaceId string `json:"adSpaceId"`
	EndTime   string `json:"endTime"`
	Status    string `json:"status"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
