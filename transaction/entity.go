package transaction

import (
	"time"

	"github.com/rezairfanwijaya/Fundraising-Website/campaign"
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
	User       user.User         // foreign key
	Campaign   campaign.Campaign // foreign key
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
