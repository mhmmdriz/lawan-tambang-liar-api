package report_file

import (
	"lawan-tambang-liar/entities"

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
