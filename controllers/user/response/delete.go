package response

import "lawan-tambang-liar/entities"

type Delete struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	ProfilePhoto string `json:"profile_photo"`
}

func DeleteFromEntitiesToResponse(user *entities.User) *Delete {
	return &Delete{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		ProfilePhoto: user.ProfilePhoto,
	}
}
