package response

import "lawan-tambang-liar/entities"

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromUseCaseToLoginResponse(admin *entities.Admin) *LoginResponse {
	return &LoginResponse{
		ID:       admin.ID,
		Name:     admin.Name,
		Username: admin.Username,
		Token:    admin.Token,
	}
}
