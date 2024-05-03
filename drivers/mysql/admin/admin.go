package admin

import (
	"errors"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"

	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{DB: db}
}

func (r *AdminRepo) CreateAccount(admin *entities.Admin) error {
	hash, _ := utils.HashPassword(admin.Password)
	(*admin).Password = hash

	if err := r.DB.Create(&admin).Error; err != nil {
		return err
	}

	return nil
}

func (r *AdminRepo) Login(admin *entities.Admin) error {
	var adminDB entities.Admin

	if err := r.DB.Where("username = ?", admin.Username).First(&adminDB).Error; err != nil {
		return errors.New("username or password is incorrect")
	}

	if !utils.CheckPasswordHash(admin.Password, adminDB.Password) {
		return errors.New("username or password is incorrect")
	}

	(*admin).ID = adminDB.ID
	(*admin).Username = adminDB.Username
	(*admin).IsSuperAdmin = adminDB.IsSuperAdmin

	return nil
}
