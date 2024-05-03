package response

import "lawan-tambang-liar/entities"

type LoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromUseCaseToLoginResponse(admin *entities.Admin) *LoginResponse {
	return &LoginResponse{
		ID:       admin.ID,
		Username: admin.Username,
		Token:    admin.Token,
	}
}
