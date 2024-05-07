package response

import "lawan-tambang-liar/entities"

type District struct {
	ID        string `json:"id"`
	RegencyID string `json:"regency_id"`
	Name      string `json:"name"`
}

func FromUseCaseToResponse(regeny *entities.District) *District {
	return &District{
		ID:        regeny.ID,
		RegencyID: regeny.RegencyID,
		Name:      regeny.Name,
	}
}
