package model

type (
	BannerResponse struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		ImageUrl    string  `json:"image_url"`
		Seq         int32   `json:"seq"`
		IsShow      bool    `json:"is_show"`
	}
)
