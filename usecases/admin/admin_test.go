package admin

import (
	"errors"
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/middlewares"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) CreateAccount(admin *entities.Admin) error {
	args := m.Called(admin)
	return args.Error(0)
}

func (m *MockAdminRepository) Login(admin *entities.Admin) error {
	args := m.Called(admin)
	return args.Error(0)
}

func (m *MockAdminRepository) GetByID(id int) (entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Admin), args.Error(1)
}

func (m *MockAdminRepository) GetAll() ([]entities.Admin, error) {
	args := m.Called()
	return args.Get(0).([]entities.Admin), args.Error(1)
}

func (m *MockAdminRepository) DeleteAccount(id int) (entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Admin), args.Error(1)
}

func (m *MockAdminRepository) ResetPassword(id int) (entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Admin), args.Error(1)
}

func (m *MockAdminRepository) ChangePassword(id int, newPassword string) (entities.Admin, error) {
	args := m.Called(id, newPassword)
	return args.Get(0).(entities.Admin), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:   "admin",
			Password:   "admin",
			Email:      "admin@example.com",
			RegencyID:  "1901",
			DistrictID: "190101",
			Address:    "Jl. Admin",
		}

		mockAdminRepo.On("CreateAccount", &admin).Return(nil)

		createdAdmin, err := adminUseCase.CreateAccount(&admin)

		assert.NoError(t, err)
		assert.Equal(t, admin, createdAdmin)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedInternalServerError", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:   "admin",
			Password:   "admin",
			Email:      "admin@example.com",
			RegencyID:  "1901",
			DistrictID: "190101",
			Address:    "Jl. Admin",
		}

		mockAdminRepo.On("CreateAccount", &admin).Return(constants.ErrInternalServerError).Once()

		createdAdmin, err := adminUseCase.CreateAccount(&admin)

		assert.Error(t, err)
		assert.Equal(t, entities.Admin{}, createdAdmin)
		assert.Equal(t, constants.ErrInternalServerError, err)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedUsernameAlreadyExist", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:   "admin",
			Password:   "admin",
			Email:      "admin@example.com",
			RegencyID:  "1901",
			DistrictID: "190101",
			Address:    "Jl. Admin",
		}

		err := errors.New("Error 1062: Duplicate entry 'admin' for key 'admins.username'")
		mockAdminRepo.On("CreateAccount", &admin).Return(err).Once()

		createdAdmin, err := adminUseCase.CreateAccount(&admin)

		assert.Error(t, err)
		assert.Equal(t, entities.Admin{}, createdAdmin)
		assert.Equal(t, constants.ErrUsernameAlreadyExist, err)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedEmailAlreadyExist", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:   "admin",
			Password:   "admin",
			Email:      "admin@example.com",
			RegencyID:  "1901",
			DistrictID: "190101",
			Address:    "Jl. Admin",
		}

		err := errors.New("Error 1062: Duplicate entry 'admin' for key 'admins.email'")
		mockAdminRepo.On("CreateAccount", &admin).Return(err).Once()

		createdAdmin, err := adminUseCase.CreateAccount(&admin)

		assert.Error(t, err)
		assert.Equal(t, entities.Admin{}, createdAdmin)
		assert.Equal(t, constants.ErrEmailAlreadyExist, err)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:  "admin",
			Password:  "admin",
			Email:     "",
			RegencyID: "",
			Address:   "",
		}

		createdAdmin, err := adminUseCase.CreateAccount(&admin)

		assert.Error(t, err)
		assert.Equal(t, entities.Admin{}, createdAdmin)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedRegencyNotFound", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:   "admin",
			Password:   "admin",
			Email:      "admin@example.com",
			RegencyID:  "9999",
			DistrictID: "190101",
			Address:    "Jl. Admin",
		}

		err := errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`lawan_tambang_liar`.`admins`, CONSTRAINT `admins_regency_id_foreign` FOREIGN KEY (`regency_id`) REFERENCES `regencies` (`id`))")
		mockAdminRepo.On("CreateAccount", &admin).Return(err).Once()

		createdAdmin, err := adminUseCase.CreateAccount(&admin)

		assert.Error(t, err)
		assert.Equal(t, entities.Admin{}, createdAdmin)
		assert.Equal(t, constants.ErrRegencyNotFound, err)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedDistrictNotFound", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:   "admin",
			Password:   "admin",
			Email:      "admin@example.com",
			RegencyID:  "1901",
			DistrictID: "9999",
			Address:    "Jl. Admin",
		}

		err := errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`lawan_tambang_liar`.`admins`, CONSTRAINT `admins_district_id_foreign` FOREIGN KEY (`district_id`) REFERENCES `districts` (`id`))")
		mockAdminRepo.On("CreateAccount", &admin).Return(err).Once()

		createdAdmin, err := adminUseCase.CreateAccount(&admin)

		assert.Error(t, err)
		assert.Equal(t, entities.Admin{}, createdAdmin)
		assert.Equal(t, constants.ErrDistrictNotFound, err)

		mockAdminRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Success Admin", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username: "admin",
			Password: "admin",
		}

		mockAdminRepo.On("Login", &admin).Return(nil)

		expectedToken := middlewares.GenerateTokenJWT(admin.ID, admin.Username, "admin")

		result, err := adminUseCase.Login(&admin)
		assert.NoError(t, err)
		assert.Equal(t, expectedToken, result.Token)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("Success Super Admin", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username:     "superadmin",
			Password:     "superadmin",
			IsSuperAdmin: true,
		}

		mockAdminRepo.On("Login", &admin).Return(nil)

		expectedToken := middlewares.GenerateTokenJWT(admin.ID, admin.Username, "super_admin")

		result, err := adminUseCase.Login(&admin)
		assert.NoError(t, err)
		assert.Equal(t, expectedToken, result.Token)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedInvalidUsernameOrPassword", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username: "admin",
			Password: "admin",
		}

		mockAdminRepo.On("Login", &admin).Return(constants.ErrInvalidUsernameOrPassword).Once()

		result, err := adminUseCase.Login(&admin)
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrInvalidUsernameOrPassword, err)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			Username: "",
			Password: "",
		}
		mockAdminRepo.On("Login", &admin).Return(constants.ErrAllFieldsMustBeFilled)

		result, err := adminUseCase.Login(&admin)
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)
	})

}

func TestGetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 1,
		}

		mockAdminRepo.On("GetByID", admin.ID).Return(admin, nil)

		result, err := adminUseCase.GetByID(admin.ID)
		assert.NoError(t, err)
		assert.Equal(t, admin, result)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedAdminNotFound", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 1,
		}

		mockAdminRepo.On("GetByID", admin.ID).Return(entities.Admin{}, constants.ErrAdminNotFound).Once()

		result, err := adminUseCase.GetByID(admin.ID)
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrAdminNotFound, err)
		mockAdminRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admins := []entities.Admin{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}

		mockAdminRepo.On("GetAll").Return(admins, nil)

		result, err := adminUseCase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, admins, result)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedInternalServerError", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)

		mockAdminRepo.On("GetAll").Return([]entities.Admin{}, constants.ErrInternalServerError).Once()

		result, err := adminUseCase.GetAll()
		assert.Equal(t, []entities.Admin{}, result)
		assert.Equal(t, constants.ErrInternalServerError, err)
		mockAdminRepo.AssertExpectations(t)
	})
}

func TestDeleteAccount(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		mockAdminRepo.On("DeleteAccount", admin.ID).Return(admin, nil).Once()

		result, err := adminUseCase.DeleteAccount(admin.ID)
		assert.NoError(t, err)
		assert.Equal(t, admin, result)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedSuperAdminCannotBeDeleted", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		mockAdminRepo.On("DeleteAccount", admin.ID).Return(entities.Admin{}, constants.ErrSuperAdminCannotBeDeleted).Once()

		result, err := adminUseCase.DeleteAccount(admin.ID)
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrSuperAdminCannotBeDeleted, err)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedAdminNotFound", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		mockAdminRepo.On("DeleteAccount", admin.ID).Return(entities.Admin{}, constants.ErrAdminNotFound).Once()

		result, err := adminUseCase.DeleteAccount(admin.ID)
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrAdminNotFound, err)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedSuperAdminCannotBeDeleted", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 1,
		}

		result, err := adminUseCase.DeleteAccount(admin.ID)
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrSuperAdminCannotBeDeleted, err)
	})
}

func TestResetPassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		mockAdminRepo.On("ResetPassword", admin.ID).Return(admin, nil).Once()

		result, err := adminUseCase.ResetPassword(admin.ID)
		assert.NoError(t, err)
		assert.Equal(t, admin, result)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedAdminNotFound", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		mockAdminRepo.On("ResetPassword", admin.ID).Return(entities.Admin{}, constants.ErrAdminNotFound).Once()

		result, err := adminUseCase.ResetPassword(admin.ID)
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrAdminNotFound, err)
		mockAdminRepo.AssertExpectations(t)
	})
}

func TestChangePassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		mockAdminRepo.On("ChangePassword", admin.ID, "newpassword").Return(admin, nil).Once()

		result, err := adminUseCase.ChangePassword(admin.ID, "newpassword")
		assert.NoError(t, err)
		assert.Equal(t, admin, result)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedAdminNotFound", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		mockAdminRepo.On("ChangePassword", admin.ID, "newpassword").Return(entities.Admin{}, constants.ErrAdminNotFound).Once()

		result, err := adminUseCase.ChangePassword(admin.ID, "newpassword")
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrAdminNotFound, err)
		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockAdminRepo := new(MockAdminRepository)
		adminUseCase := NewAdminUseCase(mockAdminRepo)
		admin := entities.Admin{
			ID: 2,
		}

		result, err := adminUseCase.ChangePassword(admin.ID, "")
		assert.Equal(t, entities.Admin{}, result)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)
	})
}
