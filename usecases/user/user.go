package user

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/middlewares"
	"strings"
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
	if user.Email == "" || user.Password == "" || user.Username == "" {
		return entities.User{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Register(user)

	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			if strings.HasSuffix(err.Error(), "email'") {
				return entities.User{}, constants.ErrEmailAlreadyExist
			} else if strings.HasSuffix(err.Error(), "username'") {
				return entities.User{}, constants.ErrUsernameAlreadyExist
			}
		} else {
			return entities.User{}, constants.ErrInternalServerError
		}
	}

	return *user, nil
}

func (u *UserUseCase) Login(user *entities.User) (entities.User, error) {
	if user.Username == "" || user.Password == "" {
		return entities.User{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Login(user)

	(*user).Token = middlewares.GenerateTokenJWT(user.ID, user.Username, "user")

	if err != nil {
		return entities.User{}, constants.ErrInvalidUsernameOrPassword
	}

	return *user, nil
}
