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
	Status                     string                      `gorm:"type:enum('verified', 'on progress', 'finished', 'rejected')"`
	CreatedAt                  time.Time                   `gorm:"autoCreateTime"`
	UpdatedAt                  time.Time                   `gorm:"autoUpdateTime"`
	DeletedAt                  gorm.DeletedAt              `gorm:"index"`
	ReportSolutionProcessFiles []ReportSolutionProcessFile `gorm:"foreignKey:ReportSolutionProcessID;references:ID"`
	Admin                      Admin                       `gorm:"foreignKey:AdminID;references:ID"`
}

type ReportSolutionProcessRepositoryInterface interface {
	Create(reportSolutionProcess *ReportSolutionProcess) error
	GetByReportID(reportID int) ([]ReportSolutionProcess, error)
	Delete(reportID int, reportSolutionProcessStatus string) (ReportSolutionProcess, error)
	Update(reportSolutionProcess ReportSolutionProcess) (ReportSolutionProcess, error)
}

type AIReportSolutionAPIInterface interface {
	GetChatCompletion(messages []map[string]string) (string, error)
}

type ReportSolutionProcessUseCaseInterface interface {
	Create(reportSolutionProcess *ReportSolutionProcess) (ReportSolutionProcess, error)
	GetByReportID(reportID int) ([]ReportSolutionProcess, error)
	Delete(reportID int, reportSolutionProcessStatus string) (ReportSolutionProcess, error)
	Update(reportSolutionProcess ReportSolutionProcess) (ReportSolutionProcess, error)
	GetMessageRecommendation(prompt string) (string, error)
}
