package entities

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type ReportSolutionProcessFile struct {
	ID                      int
	ReportSolutionProcessID int
	Path                    string
	CreatedAt               time.Time      `gorm:"autoCreateTime"`
	UpdatedAt               time.Time      `gorm:"autoUpdateTime"`
	DeletedAt               gorm.DeletedAt `gorm:"index"`
}

type ReportSolutionProcessFileRepositoryInterface interface {
	Create(reportSolutionProcessFile []*ReportSolutionProcessFile) error
	// Delete(reportSolutionProcessID int) ([]ReportSolutionProcessFile, error)
}

type ReportSolutionProcessFileGCSAPIInterface interface {
	UploadFile(files []*multipart.FileHeader) ([]string, error)
	// DeleteFile(reportSolutionProcessFile []ReportSolutionProcessFile) error
}

type ReportSolutionProcessFileUseCaseInterface interface {
	Create(files []*multipart.FileHeader, reportSolutionProcessID int) ([]ReportSolutionProcessFile, error)
	// Delete(reportSolutionProcessID int) ([]ReportSolutionProcessFile, error)
}
