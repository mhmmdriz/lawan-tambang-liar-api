package response

import "lawan-tambang-liar/entities"

type DistrictSeedingResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FromUseCaseToResponse(regeny *entities.District) *DistrictSeedingResponse {
	return &DistrictSeedingResponse{
		ID:   regeny.ID,
		Name: regeny.Name,
	}
}
