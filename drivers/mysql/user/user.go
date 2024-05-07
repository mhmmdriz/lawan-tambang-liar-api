package user

import (
	"errors"
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"time"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Register(user *entities.User) error {
	hash, _ := utils.HashPassword(user.Password)
	(*user).Password = hash

	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Login(user *entities.User) error {
	var userDB entities.User

	if err := r.DB.Where("username = ?", user.Username).First(&userDB).Error; err != nil {
		return errors.New("username or password is incorrect")
	}

	if !utils.CheckPasswordHash(user.Password, userDB.Password) {
		return errors.New("username or password is incorrect")
	}

	(*user).ID = userDB.ID
	(*user).Username = userDB.Username

	return nil
}

func (r *UserRepo) GetAll() ([]entities.User, error) {
	var users []entities.User

	if err := r.DB.Find(&users).Error; err != nil {
		return []entities.User{}, constants.ErrInternalServerError
	}

	return users, nil
}

func (r *UserRepo) GetByID(id int) (entities.User, error) {
	var user entities.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return entities.User{}, constants.ErrUserNotFound
	}

	return user, nil
}

func (r *UserRepo) Delete(id int) (entities.User, error) {
	var user entities.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return entities.User{}, constants.ErrUserNotFound
	}

	user.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := r.DB.Save(&user).Error; err != nil {
		return entities.User{}, constants.ErrInternalServerError
	}

	return user, nil
}
