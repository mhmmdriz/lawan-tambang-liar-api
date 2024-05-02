package entities

import "time"

type Admin struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	RegencyID       int       `json:"regency_id"`
	DistrictID      int       `json:"district_id"`
	Address         string    `json:"address"`
	TelephoneNumber string    `json:"telephone_number"`
	Email           string    `json:"email"`
	Username        string    `json:"username"`
	Password        string    `json:"password,omitempty"`
	ProfilePhoto    string    `gorm:"default:images/default.jpg" json:"profile_photo"`
	Token           string    `gorm:"-"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	DeletedAt       time.Time
	Regency         Regency  `gorm:"foreignKey:RegencyID"`
	District        District `gorm:"foreignKey:DistrictID"`
}

type AdminRepositoryInterface interface {
	Register(admin *Admin) error
	Login(admin *Admin) error
}

type AdminUseCaseInterface interface {
	Register(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
}
