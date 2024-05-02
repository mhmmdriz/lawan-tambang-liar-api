package response

import "lawan-tambang-liar/entities"

type RegisterResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func FromUseCaseToRegisterResponse(user *entities.User) *RegisterResponse {
	return &RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}
}
