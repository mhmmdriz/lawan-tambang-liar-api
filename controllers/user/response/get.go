package response

import "lawan-tambang-liar/entities"

type Get struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	ProfilePhoto string `json:"profile_photo"`
}

func GetFromEntitiesToResponse(user *entities.User) *Get {
	return &Get{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		ProfilePhoto: user.ProfilePhoto,
	}
}
