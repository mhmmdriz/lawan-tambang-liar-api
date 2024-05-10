package response

import (
	admin_response "lawan-tambang-liar/controllers/admin/response"
	file_response "lawan-tambang-liar/controllers/report_solution_process_file/response"
	"lawan-tambang-liar/entities"
)

type GetByReportID struct {
	ID      int                                       `json:"id"`
	Admin   admin_response.SimpleGet                  `json:"admin"`
	Status  string                                    `json:"status"`
	Message string                                    `json:"message"`
	Files   []file_response.ReportSolutionProcessFile `json:"files"`
}

func GetByReportIDFromEntitiesToResponse(reportSolutionProcess *entities.ReportSolutionProcess) *GetByReportID {
	var files []file_response.ReportSolutionProcessFile
	for _, file := range reportSolutionProcess.Files {
		files = append(files, *file_response.FromEntitiesToResponse(&file))
	}

	return &GetByReportID{
		ID:      reportSolutionProcess.ID,
		Admin:   *admin_response.SimpleGetFromEntitiesToResponse(&reportSolutionProcess.Admin),
		Status:  reportSolutionProcess.Status,
		Message: reportSolutionProcess.Message,
		Files:   files,
	}
}
