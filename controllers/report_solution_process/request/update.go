package request

import "lawan-tambang-liar/entities"

type Update struct {
	ID       int
	ReportID int
	AdminID  int
	Message  string `form:"message"`
	Status   string
}

func (r *Update) ToEntities() *entities.ReportSolutionProcess {
	return &entities.ReportSolutionProcess{
		ID:       r.ID,
		ReportID: r.ReportID,
		AdminID:  r.AdminID,
		Message:  r.Message,
		Status:   r.Status,
	}
}
