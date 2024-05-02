package response

import "lawan-tambang-liar/entities"

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromUseCaseToLoginResponse(user *entities.User) *LoginResponse {
	return &LoginResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Token:    user.Token,
	}
}
