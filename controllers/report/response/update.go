package response

import (
	"lawan-tambang-liar/controllers/report_file/response"
	"lawan-tambang-liar/entities"
)

type Update struct {
	ID          int                    `json:"id"`
	UserID      int                    `json:"user_id"`
	Title       string                 `form:"title"`
	Description string                 `form:"description"`
	RegencyID   string                 `form:"regency_id"`
	DistrictID  string                 `form:"district_id"`
	Address     string                 `form:"address"`
	Upvotes     int                    `form:"upvotes"`
	Status      string                 `form:"status"`
	Files       []*response.ReportFile `json:"files"`
}

func UpdateFromEntitiesToResponse(report *entities.Report) *Update {
	return &Update{
		ID:          report.ID,
		UserID:      report.UserID,
		Title:       report.Title,
		Description: report.Description,
		RegencyID:   report.RegencyID,
		DistrictID:  report.DistrictID,
		Address:     report.Address,
		Upvotes:     report.Upvotes,
		Status:      report.Status,
	}
}
