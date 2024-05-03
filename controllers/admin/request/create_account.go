package request

import "lawan-tambang-liar/entities"

type CreateAccount struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RegencyID  string `json:"regency_id"`
	DistrictID string `json:"district_id"`
	Address    string `json:"address"`
}

func (r *CreateAccount) ToEntities() *entities.Admin {
	return &entities.Admin{
		Username:   r.Username,
		Email:      r.Email,
		Password:   r.Password,
		RegencyID:  r.RegencyID,
		DistrictID: r.DistrictID,
		Address:    r.Address,
	}
}
