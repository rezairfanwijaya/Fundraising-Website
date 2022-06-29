package campaign

// struct input untuk mengambil data campaign berdasarkan id
// nanti id akan ditaruh di url sebagai uri
// contoh api/campaign/1
// 1 di endpoint di atas sebagai uri

type InputCampaignDetail struct {
	Id int `uri:"id" binding:"required"`
}
