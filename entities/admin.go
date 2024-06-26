package entities

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID              int    `gorm:"primaryKey"`
	Username        string `gorm:"unique"`
	RegencyID       string `gorm:"type:varchar;size:191"`
	DistrictID      string `gorm:"type:varchar;size:191"`
	Address         string
	TelephoneNumber string
	Email           string `gorm:"unique"`
	Password        string
	Token           string         `gorm:"-"`
	IsSuperAdmin    bool           `gorm:"default:false"`
	ProfilePhoto    string         `gorm:"default:profiles/default.jpg"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Regency         Regency        `gorm:"foreignKey:RegencyID"`
	District        District       `gorm:"foreignKey:DistrictID"`
}

type AdminRepositoryInterface interface {
	CreateAccount(admin *Admin) error
	Login(admin *Admin) error
	GetByID(id int) (Admin, error)
	GetAll() ([]Admin, error)
	DeleteAccount(id int) (Admin, error)
	ResetPassword(id int) (Admin, error)
	ChangePassword(id int, newPassword string) (Admin, error)
}

type AdminUseCaseInterface interface {
	CreateAccount(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
	GetByID(id int) (Admin, error)
	GetAll() ([]Admin, error)
	DeleteAccount(id int) (Admin, error)
	ResetPassword(id int) (Admin, error)
	ChangePassword(id int, newPassword string) (Admin, error)
}
