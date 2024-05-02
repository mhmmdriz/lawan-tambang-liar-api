package user

import (
	"errors"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/middlewares"
)

type UserUseCase struct {
	repository entities.UserRepositoryInterface
}

func NewUserUseCase(repository entities.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) Register(user *entities.User) (entities.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Username == "" {
		return entities.User{}, errors.New("all fields must be filled")
	}

	err := u.repository.Register(user)

	if err != nil {
		return entities.User{}, err
	}

	return *user, nil
}

func (u *UserUseCase) Login(user *entities.User) (entities.User, error) {
	if user.Username == "" || user.Password == "" {
		return entities.User{}, errors.New("all fields must be filled")
	}

	err := u.repository.Login(user)

	(*user).Token = middlewares.GenerateTokenJWT(user.ID, user.Name, "user")

	if err != nil {
		return entities.User{}, errors.New("invalid username or password")
	}

	return *user, nil
}
