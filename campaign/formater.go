package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

// struct untuk detail campaign
type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"currennt_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

// struct user untuk detail campaign
type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// struct image untuk detail campaign
type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
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

// function untuk format detail campaign
func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	campaginDetailFormatter := CampaignDetailFormatter{}
	campaginDetailFormatter.ID = campaign.Id
	campaginDetailFormatter.Name = campaign.Name
	campaginDetailFormatter.ShortDescription = campaign.ShortDescription
	campaginDetailFormatter.Description = campaign.Description
	campaginDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaginDetailFormatter.GoalAmount = campaign.GoalAmount
	campaginDetailFormatter.UserID = campaign.User.Id
	campaginDetailFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaginDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	// mengisi perks
	// pada database perks itu berupa kalimat yang dipisah oleh tanda koma, jadi kita harus mengsplit nya terlebih dahulu
	var perks []string
	data := strings.Split(campaign.Perks, ",")
	for _, perk := range data {
		perks = append(perks, strings.TrimSpace(perk)) // trimspace digunakan untuk menghapu s spasi
	}

	campaginDetailFormatter.Perks = perks

	// mengisi data user
	userCampagin := CampaignUserFormatter{}
	userCampagin.Name = campaign.User.Name
	userCampagin.ImageURL = campaign.User.AvatarFileName

	campaginDetailFormatter.User = userCampagin

	// mengisi gambar
	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaingImageFormatter := CampaignImageFormatter{}
		campaingImageFormatter.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaingImageFormatter.IsPrimary = isPrimary

		images = append(images, campaingImageFormatter)
	}

	campaginDetailFormatter.Images = images

	return campaginDetailFormatter

}
