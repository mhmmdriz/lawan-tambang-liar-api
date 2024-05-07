package report_solution_process_file

import (
	"lawan-tambang-liar/entities"
	"time"

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

func (r *ReportSolutionProcessFileRepo) Delete(reportSolutionProcessID int) ([]entities.ReportSolutionProcessFile, error) {
	var reportFiles []entities.ReportSolutionProcessFile

	// Soft delete
	if err := r.DB.Where("report_solution_process_id = ?", reportSolutionProcessID).Find(&reportFiles).Error; err != nil {
		return []entities.ReportSolutionProcessFile{}, err
	}

	for _, file := range reportFiles {
		file.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
		if err := r.DB.Save(&file).Error; err != nil {
			return []entities.ReportSolutionProcessFile{}, err
		}
	}

	return reportFiles, nil
}
