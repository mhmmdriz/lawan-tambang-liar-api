package response

import (
	"lawan-tambang-liar/controllers/report_solution_process_file/response"
	"lawan-tambang-liar/entities"
)

type Delete struct {
	ID       int                                   `json:"id"`
	ReportID int                                   `json:"report_id"`
	AdminID  int                                   `json:"admin_id"`
	Status   string                                `json:"status"`
	Message  string                                `json:"message"`
	Files    []*response.ReportSolutionProcessFile `json:"files"`
}

func DeleteFromEntitiesToResponse(report *entities.ReportSolutionProcess) *Delete {
	return &Delete{
		ID:       report.ID,
		ReportID: report.ReportID,
		AdminID:  report.AdminID,
		Status:   report.Status,
		Message:  report.Message,
	}
}
