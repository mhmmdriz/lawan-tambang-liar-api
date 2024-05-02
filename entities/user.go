package entities

import "time"

type User struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
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
}

type UserRepositoryInterface interface {
	Register(user *User) error
	Login(user *User) error
}

type UserUseCaseInterface interface {
	Register(user *User) (User, error)
	Login(user *User) (User, error)
}
