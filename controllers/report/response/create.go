package response

import (
	"lawan-tambang-liar/controllers/report_file/response"
	"lawan-tambang-liar/entities"
)

type Create struct {
	ID          int                    `json:"id"`
	UserID      int                    `json:"user_id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	RegencyID   string                 `json:"regency_id"`
	DistrictID  string                 `json:"district_id"`
	Address     string                 `json:"address"`
	Upvotes     int                    `json:"upvotes"`
	Status      string                 `json:"status"`
	Files       []*response.ReportFile `json:"files"`
}

func CreateFromEntitiesToResponse(report *entities.Report) *Create {
	return &Create{
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
