package user

import (
	"errors"
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/middlewares"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Login(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAll() ([]entities.User, error) {
	args := m.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id int) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) Delete(id int) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) ChangePassword(id int, newPassword string) (entities.User, error) {
	args := m.Called(id, newPassword)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) ResetPassword(id int) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}

func TestRegister(t *testing.T) {
	t.Run("RegisterSuccess", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userUseCase := NewUserUseCase(mockRepo)

		userEntity := &entities.User{
			Email:    "test@example.com",
			Password: "password123",
			Username: "testuser",
		}

		mockRepo.On("Register", userEntity).Return(nil)

		registeredUser, err := userUseCase.Register(userEntity)

		assert.NoError(t, err)
		assert.Equal(t, *userEntity, registeredUser)

		mockRepo.AssertExpectations(t)
	})

	t.Run("RegisterFailureInternalServerError", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userUseCase := NewUserUseCase(mockRepo)

		userEntity := &entities.User{
			Email:    "test@example.com",
			Password: "password123",
			Username: "testuser",
		}

		mockRepo.On("Register", userEntity).Return(constants.ErrInternalServerError)

		registeredUser, err := userUseCase.Register(userEntity)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.User{}, registeredUser)

		mockRepo.AssertExpectations(t)
	})

	t.Run("RegisterFailureEmailAlreadyExist", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userUseCase := NewUserUseCase(mockRepo)

		userEntity := &entities.User{
			Email:    "test@example.com",
			Password: "password123",
			Username: "testuser",
		}

		err := errors.New("Error 1062: Duplicate entry 'test@example.com' for key 'users.email'")
		mockRepo.On("Register", userEntity).Return(err)

		registeredUser, err := userUseCase.Register(userEntity)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrEmailAlreadyExist, err)
		assert.Equal(t, entities.User{}, registeredUser)

		mockRepo.AssertExpectations(t)
	})

	t.Run("RegisterFailureUsernameAlreadyExist", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userUseCase := NewUserUseCase(mockRepo)

		userEntity := &entities.User{
			Email:    "test@example.com",
			Password: "password123",
			Username: "testuser",
		}

		err := errors.New("Error 1062: Duplicate entry 'testuser' for key 'users.username'")
		mockRepo.On("Register", userEntity).Return(err)

		registeredUser, err := userUseCase.Register(userEntity)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUsernameAlreadyExist, err)
		assert.Equal(t, entities.User{}, registeredUser)

		mockRepo.AssertExpectations(t)
	})

	t.Run("RegisterFailureAllFieldsMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		userUseCase := NewUserUseCase(mockRepo)

		userEntity := &entities.User{
			Email:    "",
			Password: "password123",
			Username: "testuser",
		}

		expectedError := constants.ErrAllFieldsMustBeFilled

		registeredUser, err := userUseCase.Register(userEntity)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.User{}, registeredUser)

		mockRepo.AssertNotCalled(t, "Register")
	})
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUserUseCase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockUser := &entities.User{Username: "testuser", Password: "password123"}
		mockRepo.On("Login", mockUser).Return(nil)

		expectedToken := middlewares.GenerateTokenJWT(mockUser.ID, mockUser.Username, "user")

		result, err := userUseCase.Login(mockUser)
		assert.NoError(t, err)
		assert.Equal(t, expectedToken, result.Token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InvalidUsernameOrPassword", func(t *testing.T) {
		mockUser := &entities.User{Username: "testuser", Password: "password123"}
		mockRepo.On("Login", mockUser).Return(constants.ErrInvalidUsernameOrPassword)

		_, err := userUseCase.Login(mockUser)
		assert.Equal(t, constants.ErrInvalidUsernameOrPassword, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("AllFieldsMustBeFilled", func(t *testing.T) {
		mockUser := &entities.User{Username: "", Password: "password"}
		mockRepo.On("Login", mockUser).Return(constants.ErrAllFieldsMustBeFilled)

		_, err := userUseCase.Login(mockUser)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)
	})
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUserUseCase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockUsers := []entities.User{
			{ID: 1, Username: "testuser1"},
			{ID: 2, Username: "testuser2"},
		}
		mockRepo.On("GetAll").Return(mockUsers, nil).Once()

		result, err := userUseCase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, mockUsers, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		mockRepo.On("GetAll").Return([]entities.User{}, constants.ErrInternalServerError).Once()

		result, err := userUseCase.GetAll()
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, []entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUserUseCase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("GetByID", mockUser.ID).Return(mockUser, nil).Once()

		result, err := userUseCase.GetByID(mockUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockUser, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("GetByID", mockUser.ID).Return(entities.User{}, constants.ErrUserNotFound).Once()

		result, err := userUseCase.GetByID(mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserNotFound, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("GetByID", mockUser.ID).Return(entities.User{}, constants.ErrInternalServerError).Once()

		result, err := userUseCase.GetByID(mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUserUseCase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("Delete", mockUser.ID).Return(mockUser, nil).Once()

		result, err := userUseCase.Delete(mockUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockUser, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("Delete", mockUser.ID).Return(entities.User{}, constants.ErrUserNotFound).Once()

		result, err := userUseCase.Delete(mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserNotFound, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("Delete", mockUser.ID).Return(entities.User{}, constants.ErrInternalServerError).Once()

		result, err := userUseCase.Delete(mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestChangePassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUserUseCase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		newPassword := "newpassword123"
		mockRepo.On("ChangePassword", mockUser.ID, newPassword).Return(mockUser, nil).Once()

		result, err := userUseCase.ChangePassword(mockUser.ID, newPassword)
		assert.NoError(t, err)
		assert.Equal(t, mockUser, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("AllFieldsMustBeFilled", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		newPassword := ""
		// mockRepo.On("ChangePassword", mockUser.ID, newPassword).Return(entities.User{}, constants.ErrAllFieldsMustBeFilled).Once()

		result, err := userUseCase.ChangePassword(mockUser.ID, newPassword)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		newPassword := "newpassword123"
		mockRepo.On("ChangePassword", mockUser.ID, newPassword).Return(entities.User{}, constants.ErrInternalServerError).Once()

		result, err := userUseCase.ChangePassword(mockUser.ID, newPassword)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		newPassword := "newpassword123"
		mockRepo.On("ChangePassword", mockUser.ID, newPassword).Return(entities.User{}, constants.ErrUserNotFound).Once()

		result, err := userUseCase.ChangePassword(mockUser.ID, newPassword)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserNotFound, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestResetPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUserUseCase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("ResetPassword", mockUser.ID).Return(mockUser, nil).Once()

		result, err := userUseCase.ResetPassword(mockUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockUser, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InternalServerError", func(t *testing.T) {
		mockUser := entities.User{ID: 1, Username: "testuser"}
		mockRepo.On("ResetPassword", mockUser.ID).Return(entities.User{}, constants.ErrInternalServerError).Once()

		result, err := userUseCase.ResetPassword(mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInternalServerError, err)
		assert.Equal(t, entities.User{}, result)
		mockRepo.AssertExpectations(t)
	})
}
