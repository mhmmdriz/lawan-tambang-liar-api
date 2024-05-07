package response

import "lawan-tambang-liar/entities"

type CreateAccount struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func CreateAccountFromEntitiesToResponse(admin *entities.Admin) *CreateAccount {
	return &CreateAccount{
		ID:       admin.ID,
		Email:    admin.Email,
		Username: admin.Username,
	}
}
