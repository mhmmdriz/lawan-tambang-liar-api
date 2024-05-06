package response

import "lawan-tambang-liar/entities"

type ReportFile struct {
	ID       int    `json:"id"`
	ReportID int    `json:"report_id"`
	Path     string `json:"path"`
}

func FromEntitiesToResponse(file *entities.ReportFile) *ReportFile {
	return &ReportFile{
		ID:       file.ID,
		ReportID: file.ReportID,
		Path:     file.Path,
	}
}
