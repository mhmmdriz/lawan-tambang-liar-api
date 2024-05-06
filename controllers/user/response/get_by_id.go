package response

import "lawan-tambang-liar/entities"

type GetByID struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	ProfilePhoto string `json:"profile_photo"`
}

func GetByIDFromEntitiesToResponse(user *entities.User) *GetByID {
	return &GetByID{
		ID:           user.ID,
		Username:     user.Username,
		ProfilePhoto: user.ProfilePhoto,
	}
}
