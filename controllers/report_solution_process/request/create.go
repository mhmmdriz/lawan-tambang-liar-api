package request

import "lawan-tambang-liar/entities"

type Create struct {
	ReportID int
	AdminID  int
	Message  string `form:"message"`
	Status   string
}

func (r *Create) ToEntities() *entities.ReportSolutionProcess {
	return &entities.ReportSolutionProcess{
		ReportID: r.ReportID,
		AdminID:  r.AdminID,
		Message:  r.Message,
		Status:   r.Status,
	}
}
