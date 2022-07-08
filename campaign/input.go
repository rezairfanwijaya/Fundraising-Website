package campaign

import user "github.com/rezairfanwijaya/Fundraising-Website/users"

// struct input untuk mengambil data campaign berdasarkan id
// nanti id akan ditaruh di url sebagai uri
// contoh api/campaign/1
// 1 di endpoint di atas sebagai uri

type InputCampaignDetail struct {
	Id int `uri:"id" binding:"required"`
}

// struct input create campaign
type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

// struct input upload campaign image
type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"` // jangan set binding required ketika tipe bool, karena pas kita assign false akan memunculkan error required tag
	User       user.User
}
