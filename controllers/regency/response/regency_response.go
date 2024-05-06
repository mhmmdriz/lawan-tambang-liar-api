package response

import "lawan-tambang-liar/entities"

type Regency struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FromUseCaseToResponse(regeny *entities.Regency) *Regency {
	return &Regency{
		ID:   regeny.ID,
		Name: regeny.Name,
	}
}
