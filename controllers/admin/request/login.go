package request

import "lawan-tambang-liar/entities"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Login) ToEntities() *entities.Admin {
	return &entities.Admin{
		Username: r.Username,
		Password: r.Password,
	}
}
