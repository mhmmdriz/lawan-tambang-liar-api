package response

import (
	"lawan-tambang-liar/entities"
)

type SimpleGet struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	RegencyID       string `json:"regency_id"`
	DistrictID      string `json:"district_id"`
	Address         string `json:"address"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
	IsSuperAdmin    bool   `json:"is_super_admin"`
	ProfilePhoto    string `json:"profile_photo"`
}

func SimpleGetFromEntitiesToResponse(admin *entities.Admin) *SimpleGet {
	return &SimpleGet{
		ID:              admin.ID,
		Username:        admin.Username,
		RegencyID:       admin.RegencyID,
		DistrictID:      admin.DistrictID,
		Address:         admin.Address,
		TelephoneNumber: admin.TelephoneNumber,
		Email:           admin.Email,
		IsSuperAdmin:    admin.IsSuperAdmin,
		ProfilePhoto:    admin.ProfilePhoto,
	}
}
