package report

import (
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

type ReportRepo struct {
	DB *gorm.DB
}

func NewReportRepo(db *gorm.DB) *ReportRepo {
	return &ReportRepo{
		DB: db,
	}
}

func (r *ReportRepo) Create(report *entities.Report) error {
	if err := r.DB.Create(&report).Error; err != nil {
		return err
	}
	return nil
}
