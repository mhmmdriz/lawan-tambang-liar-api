package admin

import (
	"errors"
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/middlewares"
	"strings"
)

type AdminUseCase struct {
	repository entities.AdminRepositoryInterface
}

func NewAdminUseCase(repository entities.AdminRepositoryInterface) *AdminUseCase {
	return &AdminUseCase{
		repository: repository,
	}
}

func (u *AdminUseCase) CreateAccount(admin *entities.Admin) (entities.Admin, error) {
	if admin.Email == "" || admin.Password == "" || admin.Username == "" || admin.RegencyID == "" || admin.DistrictID == "" || admin.Address == "" {
		return entities.Admin{}, errors.New("all fields must be filled")
	}

	err := u.repository.CreateAccount(admin)

	if err != nil {
		return entities.Admin{}, err
	}

	return *admin, nil
}

func (u *AdminUseCase) Login(admin *entities.Admin) (entities.Admin, error) {
	if admin.Username == "" || admin.Password == "" {
		return entities.Admin{}, errors.New("all fields must be filled")
	}

	err := u.repository.Login(admin)

	if admin.IsSuperAdmin {
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Username, "super_admin")
	} else {
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Username, "admin")
	}

	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			if strings.HasSuffix(err.Error(), "email'") {
				return entities.Admin{}, constants.ErrEmailAlreadyExist
			} else if strings.HasSuffix(err.Error(), "username'") {
				return entities.Admin{}, constants.ErrUsernameAlreadyExist
			}
		} else {
			return entities.Admin{}, constants.ErrInternalServerError
		}
	}

	return *admin, nil
}
