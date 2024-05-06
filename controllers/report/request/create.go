package request

import "lawan-tambang-liar/entities"

type Create struct {
	UserID      int    `form:"user_id"`
	Title       string `form:"title"`
	Description string `form:"description"`
	RegencyID   string `form:"regency_id"`
	DistrictID  string `form:"district_id"`
	Address     string `form:"address"`
}

func (r *Create) ToEntities() *entities.Report {
	return &entities.Report{
		UserID:      r.UserID,
		Title:       r.Title,
		Description: r.Description,
		RegencyID:   r.RegencyID,
		DistrictID:  r.DistrictID,
		Address:     r.Address,
	}
}
