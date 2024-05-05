package entities

import (
	"time"

	"gorm.io/gorm"
)

type ReportSolutionProcessFile struct {
	ID        int
	ReportID  int
	Path      string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
