package entities

import (
	"time"

	"gorm.io/gorm"
)

type ReportSolutionProcess struct {
	ID                         int `gorm:"primaryKey"`
	ReportID                   int
	AdminID                    int
	Message                    string
	Status                     string                      `gorm:"type:enum('verified', 'on progress', 'done', 'rejected')"`
	CreatedAt                  time.Time                   `gorm:"autoCreateTime"`
	UpdatedAt                  time.Time                   `gorm:"autoUpdateTime"`
	DeletedAt                  gorm.DeletedAt              `gorm:"index"`
	ReportSolutionProcessFiles []ReportSolutionProcessFile `gorm:"foreignKey:ReportSolutionProcessID;references:ID"`
}

type ReportSolutionProcessRepositoryInterface interface {
	Create(reportSolutionProcess *ReportSolutionProcess) error
	// GetByReportID(reportID int) ([]ReportSolutionProcess, error)
	// Delete(reportSolutionProcessID int) error
	// Update(reportSolutionProcess *ReportSolutionProcess) error
}

type ReportSolutionProcessUseCaseInterface interface {
	Create(reportSolutionProcess *ReportSolutionProcess) (ReportSolutionProcess, error)
	// GetByReportID(reportID int) ([]ReportSolutionProcess, error)
	// Delete(reportSolutionProcessID int) error
	// Update(reportSolutionProcess *ReportSolutionProcess) error
}
