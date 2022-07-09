package transaction

import (
	"time"

	user "github.com/rezairfanwijaya/Fundraising-Website/users"
)

// struct models table transaction
type Transaction struct {
	Id         int
	CampaignID int
	UserId     int
	Amount     int
	Status     string
	Code       string
	User       user.User // foreign key
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
