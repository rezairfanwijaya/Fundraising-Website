package user

// bikin struct sesuai yang diminta sama helper
type UserFormat struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Occupation string `json:"occupation"`
	ImageURL   string `json:"image_url"`
}

func UserFormatter(user User, token string) UserFormat {
	userFormat := UserFormat{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Token:      token,
		Occupation: user.Occupation,
		ImageURL:   user.AvatarFileName,
	}

	return userFormat
}
