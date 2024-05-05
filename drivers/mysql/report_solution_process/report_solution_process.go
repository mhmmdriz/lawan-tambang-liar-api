package report_solution_process

import (
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

type ReportSolutionProcessRepo struct {
	DB *gorm.DB
}

func NewReportSolutionProcessRepo(db *gorm.DB) *ReportSolutionProcessRepo {
	return &ReportSolutionProcessRepo{
		DB: db,
	}
}

func (r *ReportSolutionProcessRepo) Create(reportSolutionProcess *entities.ReportSolutionProcess) error {
	if err := r.DB.Create(&reportSolutionProcess).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReportSolutionProcessRepo) GetByReportID(reportID int) ([]entities.ReportSolutionProcess, error) {
	var reportSolutionProcesses []entities.ReportSolutionProcess

	if err := r.DB.Where("report_id = ?", reportID).Find(&reportSolutionProcesses).Error; err != nil {
		return nil, err
	}

	return reportSolutionProcesses, nil
}
