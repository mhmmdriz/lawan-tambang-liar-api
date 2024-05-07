package response

import (
	district_response "lawan-tambang-liar/controllers/district/response"
	regency_response "lawan-tambang-liar/controllers/regency/response"
	file_response "lawan-tambang-liar/controllers/report_file/response"
	user_response "lawan-tambang-liar/controllers/user/response"
	"lawan-tambang-liar/entities"
)

type GetPaginate struct {
	ID          int                        `json:"id"`
	User        user_response.GetByID      `json:"user"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Regency     regency_response.Regency   `json:"regency"`
	District    district_response.District `json:"district"`
	Address     string                     `json:"address"`
	Upvotes     int                        `json:"upvotes"`
	Status      string                     `json:"status"`
	Files       []file_response.ReportFile `json:"files"`
}

func GetPaginateFromEntitiesToResponse(report *entities.Report) *GetPaginate {
	var files []file_response.ReportFile
	for _, file := range report.Files {
		files = append(files, *file_response.FromEntitiesToResponse(&file))
	}

	return &GetPaginate{
		ID:          report.ID,
		User:        *user_response.GetByIDFromEntitiesToResponse(&report.User),
		Title:       report.Title,
		Description: report.Description,
		Regency:     *regency_response.FromUseCaseToResponse(&report.Regency),
		District:    *district_response.FromUseCaseToResponse(&report.District),
		Address:     report.Address,
		Upvotes:     report.Upvotes,
		Status:      report.Status,
		Files:       files,
	}
}
