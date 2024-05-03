package request

import "lawan-tambang-liar/entities"

type Register struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Register) ToEntities() *entities.User {
	return &entities.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
	}
}
