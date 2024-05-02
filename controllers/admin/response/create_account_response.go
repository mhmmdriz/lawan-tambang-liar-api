package response

import "lawan-tambang-liar/entities"

type CreateAccountResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func FromUseCaseToCreateAccountResponse(admin *entities.Admin) *CreateAccountResponse {
	return &CreateAccountResponse{
		ID:       admin.ID,
		Name:     admin.Name,
		Email:    admin.Email,
		Username: admin.Username,
	}
}
