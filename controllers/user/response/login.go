package response

import "lawan-tambang-liar/entities"

type Login struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func LoginFromEntitiesToResponse(user *entities.User) *Login {
	return &Login{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}
}
