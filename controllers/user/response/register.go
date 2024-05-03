package response

import "lawan-tambang-liar/entities"

type Register struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func RegisterFromEntitiesToResponse(user *entities.User) *Register {
	return &Register{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}
