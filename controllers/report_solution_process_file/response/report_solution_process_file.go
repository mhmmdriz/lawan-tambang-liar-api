package response

import "lawan-tambang-liar/entities"

type ReportSolutionProcessFile struct {
	ID                      int    `json:"id"`
	ReportSolutionProcessID int    `json:"report_solution_process_id"`
	Path                    string `json:"path"`
}

func FromEntitiesToResponse(file *entities.ReportSolutionProcessFile) *ReportSolutionProcessFile {
	return &ReportSolutionProcessFile{
		ID:                      file.ID,
		ReportSolutionProcessID: file.ReportSolutionProcessID,
		Path:                    file.Path,
	}
}
