package response

import (
	"lawan-tambang-liar/controllers/report_solution_process_file/response"
	"lawan-tambang-liar/entities"
)

type Create struct {
	ID       int                                   `json:"id"`
	ReportID int                                   `json:"report_id"`
	AdminID  int                                   `json:"admin_id"`
	Message  string                                `json:"message"`
	Status   string                                `json:"status"`
	Files    []*response.ReportSolutionProcessFile `json:"file"`
}

func CreateFromEntitiesToResponse(reportSolutionProcess *entities.ReportSolutionProcess) *Create {
	return &Create{
		ID:       reportSolutionProcess.ID,
		ReportID: reportSolutionProcess.ReportID,
		AdminID:  reportSolutionProcess.AdminID,
		Message:  reportSolutionProcess.Message,
		Status:   reportSolutionProcess.Status,
	}
}
