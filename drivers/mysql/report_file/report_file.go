package report_file

import (
	"lawan-tambang-liar/entities"
	"time"

	"gorm.io/gorm"
)

type ReportFileRepo struct {
	DB *gorm.DB
}

func NewReportFileRepo(db *gorm.DB) *ReportFileRepo {
	return &ReportFileRepo{
		DB: db,
	}
}

func (r *ReportFileRepo) Create(reportFile []*entities.ReportFile) error {
	if err := r.DB.CreateInBatches(&reportFile, len(reportFile)).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReportFileRepo) Delete(report_id int) ([]entities.ReportFile, error) {
	var reportFile []entities.ReportFile

	// Soft delete
	if err := r.DB.Where("report_id = ?", report_id).Find(&reportFile).Error; err != nil {
		return []entities.ReportFile{}, err
	}

	for _, file := range reportFile {
		file.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
		if err := r.DB.Save(&file).Error; err != nil {
			return []entities.ReportFile{}, err
		}
	}

	return reportFile, nil
}
