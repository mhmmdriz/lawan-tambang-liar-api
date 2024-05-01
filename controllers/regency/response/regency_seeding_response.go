package response

import "lawan-tambang-liar/entities"

type RegencySeedingResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FromUseCaseToResponse(regeny *entities.Regency) *RegencySeedingResponse {
	return &RegencySeedingResponse{
		ID:   regeny.ID,
		Name: regeny.Name,
	}
}
