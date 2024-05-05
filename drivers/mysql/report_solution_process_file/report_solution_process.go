package report_solution_process_file

import (
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

type ReportSolutionProcessFileRepo struct {
	DB *gorm.DB
}

func NewReportSolutionProcessFileRepo(db *gorm.DB) *ReportSolutionProcessFileRepo {
	return &ReportSolutionProcessFileRepo{
		DB: db,
	}
}

func (r *ReportSolutionProcessFileRepo) Create(reportFile []*entities.ReportSolutionProcessFile) error {
	if err := r.DB.CreateInBatches(&reportFile, len(reportFile)).Error; err != nil {
		return err
	}
	return nil
}
