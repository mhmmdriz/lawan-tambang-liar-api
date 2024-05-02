package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name"`
	Address         string         `json:"address"`
	TelephoneNumber string         `json:"telephone_number"`
	Email           string         `json:"email" gorm:"unique"`
	Username        string         `json:"username" gorm:"unique"`
	Password        string         `json:"password"`
	ProfilePhoto    string         `gorm:"default:images/default.jpg" json:"profile_photo"`
	Token           string         `gorm:"-"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type UserRepositoryInterface interface {
	Register(user *User) error
	Login(user *User) error
}

type UserUseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
}
