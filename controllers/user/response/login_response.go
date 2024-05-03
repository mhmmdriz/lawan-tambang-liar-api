package response

import "lawan-tambang-liar/entities"

type LoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromUseCaseToLoginResponse(user *entities.User) *LoginResponse {
	return &LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}
}
