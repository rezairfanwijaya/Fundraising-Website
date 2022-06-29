package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	Slug             string `json:"slug"`
	CurrentAmount    int    `json:"current_amount"`
}

// function untuk format campaign
func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.Id
	campaignFormatter.UserID = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

// function untuk format jamak campaign
// kita perlu function ini karena kita tahu bahwa balikan dari s.repository.FindByUserId(userID) di file campaign.go folder handler function GetCampaigns ini adalah slice of campaign
func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormat := []CampaignFormatter{}

	// loop
	for _, campaign := range campaigns {
		// lakukan format terhadap single campaign
		campaignFormat := FormatCampaign(campaign)
		// append
		campaignsFormat = append(campaignsFormat, campaignFormat)
	}

	return campaignsFormat
}
