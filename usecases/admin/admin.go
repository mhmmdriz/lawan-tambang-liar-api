package admin

import (
	"errors"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/middlewares"
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
	if admin.Name == "" || admin.Email == "" || admin.Password == "" || admin.Username == "" || admin.RegencyID == "" || admin.DistrictID == "" || admin.Address == "" {
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
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Name, "super_admin")
	} else {
		(*admin).Token = middlewares.GenerateTokenJWT(admin.ID, admin.Name, "admin")
	}

	if err != nil {
		return entities.Admin{}, errors.New("invalid username or password")
	}

	return *admin, nil
}
