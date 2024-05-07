package response

import (
	district_response "lawan-tambang-liar/controllers/district/response"
	regency_response "lawan-tambang-liar/controllers/regency/response"
	"lawan-tambang-liar/entities"
)

type Get struct {
	ID              int                        `json:"id"`
	Username        string                     `json:"username"`
	Regency         regency_response.Regency   `json:"regency"`
	District        district_response.District `json:"district"`
	Address         string                     `json:"address"`
	TelephoneNumber string                     `json:"telephone_number"`
	Email           string                     `json:"email"`
	IsSuperAdmin    bool                       `json:"is_super_admin"`
	ProfilePhoto    string                     `json:"profile_photo"`
}

func GetFromEntitiesToResponse(admin *entities.Admin) *Get {
	return &Get{
		ID:              admin.ID,
		Username:        admin.Username,
		Regency:         *regency_response.FromUseCaseToResponse(&admin.Regency),
		District:        *district_response.FromUseCaseToResponse(&admin.District),
		Address:         admin.Address,
		TelephoneNumber: admin.TelephoneNumber,
		Email:           admin.Email,
		IsSuperAdmin:    admin.IsSuperAdmin,
		ProfilePhoto:    admin.ProfilePhoto,
	}
}
