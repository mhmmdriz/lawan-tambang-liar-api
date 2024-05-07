package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int    `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	ProfilePhoto string         `gorm:"default:profiles/default.jpg"`
	Token        string         `gorm:"-"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type UserRepositoryInterface interface {
	Register(user *User) error
	Login(user *User) error
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Delete(id int) (User, error)
	// ChangePassword(id int, newPassword string) (User, error)
	// ResetPassword(id int) (User, error)
}

type UserUseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Delete(id int) (User, error)
	// ChangePassword(id int, newPassword string) (User, error)
	// ResetPassword(id int) (User, error)
}
