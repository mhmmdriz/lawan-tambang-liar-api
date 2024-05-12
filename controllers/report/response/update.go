package response

import (
	district_response "lawan-tambang-liar/controllers/district/response"
	regency_response "lawan-tambang-liar/controllers/regency/response"
	report_file_response "lawan-tambang-liar/controllers/report_file/response"
	user_response "lawan-tambang-liar/controllers/user/response"
	"lawan-tambang-liar/entities"
)

type Update struct {
	ID          int                                `json:"id"`
	User        *user_response.GetByID             `json:"user"`
	Title       string                             `json:"title"`
	Description string                             `json:"description"`
	Regency     *regency_response.Regency          `json:"regency"`
	District    *district_response.District        `json:"district"`
	Address     string                             `json:"address"`
	Upvotes     int                                `json:"upvotes"`
	Status      string                             `json:"status"`
	Files       []*report_file_response.ReportFile `json:"files"`
}

func UpdateFromEntitiesToResponse(report *entities.Report) *Update {
	return &Update{
		ID:          report.ID,
		User:        user_response.GetByIDFromEntitiesToResponse(&report.User),
		Title:       report.Title,
		Description: report.Description,
		Regency:     regency_response.FromUseCaseToResponse(&report.Regency),
		District:    district_response.FromUseCaseToResponse(&report.District),
		Address:     report.Address,
		Upvotes:     report.Upvotes,
		Status:      report.Status,
	}
}
