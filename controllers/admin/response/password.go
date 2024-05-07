package response

import "lawan-tambang-liar/entities"

type Password struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func PasswordFromEntitiesToResponse(admin *entities.Admin) *Password {
	return &Password{
		ID:       admin.ID,
		Email:    admin.Email,
		Username: admin.Username,
		Password: admin.Password,
	}
}
