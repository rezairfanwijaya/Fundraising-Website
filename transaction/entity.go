package transaction

import "time"

// struct models table transaction
type Transaction struct {
	Id         int
	CampaignID int
	UserId     int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
