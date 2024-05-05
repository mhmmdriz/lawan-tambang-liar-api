package entities

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type ReportFile struct {
	ID        int
	ReportID  int
	Path      string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ReportFileRepositoryInterface interface {
	Create(reportFile []*ReportFile) error
	Delete(report_id int) ([]ReportFile, error)
}

type ReportFileGCSAPIInterface interface {
	UploadFile(files []*multipart.FileHeader) ([]string, error)
	DeleteFile(filePaths []string) error
}

type ReportFileUseCaseInterface interface {
	Create(files []*multipart.FileHeader, report_id int) ([]ReportFile, error)
	Delete(report_id int) ([]ReportFile, error)
}
