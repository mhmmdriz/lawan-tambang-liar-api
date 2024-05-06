package response

import "lawan-tambang-liar/entities"

type GetByReportID struct {
	ID       int    `json:"id"`
	ReportID int    `json:"report_id"`
	AdminID  int    `json:"admin_id"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

func GetByReportIDFromEntitiesToResponse(reportSolutionProcess *entities.ReportSolutionProcess) *GetByReportID {
	return &GetByReportID{
		ID:       reportSolutionProcess.ID,
		ReportID: reportSolutionProcess.ReportID,
		AdminID:  reportSolutionProcess.AdminID,
		Status:   reportSolutionProcess.Status,
		Message:  reportSolutionProcess.Message,
	}
}
