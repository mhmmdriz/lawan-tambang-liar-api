package response

import "lawan-tambang-liar/entities"

type ReportFile struct {
	ReportID int    `json:"report_id"`
	Path     string `json:"path"`
}

func CreateFromEntitiesToResponse(file *entities.ReportFile) *ReportFile {
	return &ReportFile{
		ReportID: file.ReportID,
		Path:     file.Path,
	}
}
