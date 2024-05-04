package entities

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID          int
	UserID      int
	Title       string
	Description string
	RegencyID   string `gorm:"type:varchar;size:191"`
	DistrictID  string `gorm:"type:varchar;size:191"`
	Address     string
	Status      string `gorm:"default:'verification';type:enum('verification', 'on progress', 'done', 'rejected')"`
	Upvotes     int
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User        User           `gorm:"foreignKey:UserID;references:ID"`
	Regency     Regency        `gorm:"foreignKey:RegencyID;references:ID"`
	District    District       `gorm:"foreignKey:DistrictID;references:ID"`
	Files       []ReportFile   `gorm:"foreignKey:ReportID;references:ID"`
}

type ReportRepositoryInterface interface {
	Create(report *Report) error
	// GetPaginate(limit int, page int) ([]Report, error)
	// GetByID(id int) (Report, error)
	// Update(id int) (Report, error)
	// Delete(id int) (Report, error)
}

type ReportUseCaseInterface interface {
	Create(report *Report) (Report, error)
	// GetPaginate(limit int, page int) ([]Report, error)
	// GetByID(id int) (Report, error)
	// Update(id int) (Report, error)
	// Delete(id int) (Report, error)
}
