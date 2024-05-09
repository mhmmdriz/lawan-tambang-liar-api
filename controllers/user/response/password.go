package response

import "lawan-tambang-liar/entities"

type Password struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PasswordFromEntitiesToResponse(user *entities.User) *Password {
	return &Password{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
