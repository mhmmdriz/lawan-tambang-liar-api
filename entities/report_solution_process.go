package entities

import (
	"time"

	"gorm.io/gorm"
)

type ReportSolutionProcess struct {
	ID        int `gorm:"primaryKey"`
	ReportID  int
	AdminID   int
	Message   string
	Status    string         `gorm:"type:enum('verification', 'on progress', 'done', 'rejected')"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
