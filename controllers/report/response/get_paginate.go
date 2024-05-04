package response

import (
	"lawan-tambang-liar/entities"
)

type GetPaginate struct {
	ID          int      `json:"id"`
	User        string   `json:"user"`
	Title       string   `form:"title"`
	Description string   `form:"description"`
	Regency     string   `form:"regency_id"`
	District    string   `form:"district_id"`
	Address     string   `form:"address"`
	Upvotes     int      `form:"upvotes"`
	Status      string   `form:"status"`
	Files       []string `json:"files"`
}

func GetPaginateFromEntitiesToResponse(report *entities.Report) *GetPaginate {
	return &GetPaginate{
		ID:          report.ID,
		Title:       report.Title,
		Description: report.Description,
		Address:     report.Address,
		Upvotes:     report.Upvotes,
		Status:      report.Status,
	}
}
