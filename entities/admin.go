package entities

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name"`
	RegencyID       string         `json:"regency_id" gorm:"type:varchar;size:191"`
	DistrictID      string         `json:"district_id" gorm:"type:varchar;size:191"`
	Address         string         `json:"address"`
	TelephoneNumber string         `json:"telephone_number"`
	Email           string         `json:"email" gorm:"unique"`
	Username        string         `json:"username" gorm:"unique"`
	Password        string         `json:"password"`
	ProfilePhoto    string         `gorm:"default:images/default.jpg" json:"profile_photo"`
	Token           string         `gorm:"-"`
	IsSuperAdmin    bool           `json:"is_super_admin"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Regency         Regency        `gorm:"foreignKey:RegencyID"`
	District        District       `gorm:"foreignKey:DistrictID"`
}

type AdminRepositoryInterface interface {
	CreateAccount(admin *Admin) error
	Login(admin *Admin) error
}

type AdminUseCaseInterface interface {
	CreateAccount(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
}
