package regency

import (
	"errors"
	"lawan-tambang-liar/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRegencyRepository struct {
	mock.Mock
}

func (m *MockRegencyRepository) AddRegenciesFromAPI(regencies []entities.Regency) error {
	args := m.Called(regencies)
	return args.Error(0)
}

func (m *MockRegencyRepository) GetRegencyIDs() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockRegencyRepository) GetAll() ([]entities.Regency, error) {
	args := m.Called()
	return args.Get(0).([]entities.Regency), args.Error(1)
}

func (m *MockRegencyRepository) GetByID(id string) (entities.Regency, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Regency), args.Error(1)
}

type MockRegencyAPI struct {
	mock.Mock
}

func (m *MockRegencyAPI) GetRegenciesDataFromAPI() ([]entities.Regency, error) {
	args := m.Called()
	return args.Get(0).([]entities.Regency), args.Error(1)
}

func TestSeedRegencyDBFromAPI(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedRegencies := []entities.Regency{
			{ID: "1", Name: "Regency 1"},
			{ID: "2", Name: "Regency 2"},
		}

		mockAPI.On("GetRegenciesDataFromAPI").Return(expectedRegencies, nil)
		mockRepo.On("AddRegenciesFromAPI", expectedRegencies).Return(nil)

		resultRegencies, err := regencyUseCase.SeedRegencyDBFromAPI()

		assert.NoError(t, err)
		assert.Equal(t, expectedRegencies, resultRegencies)
	})

	t.Run("FailedRepositoryError", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedRegencies := []entities.Regency{
			{ID: "1", Name: "Regency 1"},
			{ID: "2", Name: "Regency 2"},
		}

		mockAPI.On("GetRegenciesDataFromAPI").Return(expectedRegencies, nil)
		mockRepo.On("AddRegenciesFromAPI", expectedRegencies).Return(errors.New("error"))

		resultRegencies, err := regencyUseCase.SeedRegencyDBFromAPI()

		assert.Error(t, err)
		assert.Nil(t, resultRegencies)
	})

	t.Run("FailedAPIError", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedError := errors.New("API error")

		mockAPI.On("GetRegenciesDataFromAPI").Return([]entities.Regency{}, expectedError)

		resultRegencies, err := regencyUseCase.SeedRegencyDBFromAPI()

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, resultRegencies)
	})
}

func TestGetRegencyIDs(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedRegencyIDs := []string{"1", "2"}

		mockRepo.On("GetRegencyIDs").Return(expectedRegencyIDs, nil)

		resultRegencyIDs, err := regencyUseCase.GetRegencyIDs()

		assert.NoError(t, err)
		assert.Equal(t, expectedRegencyIDs, resultRegencyIDs)
	})

	t.Run("FailedRepositoryError", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedError := errors.New("error")

		mockRepo.On("GetRegencyIDs").Return([]string{}, expectedError)

		resultRegencyIDs, err := regencyUseCase.GetRegencyIDs()

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, resultRegencyIDs)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedRegencies := []entities.Regency{
			{ID: "1", Name: "Regency 1"},
			{ID: "2", Name: "Regency 2"},
		}

		mockRepo.On("GetAll").Return(expectedRegencies, nil)

		resultRegencies, err := regencyUseCase.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, expectedRegencies, resultRegencies)
	})

	t.Run("FailedRepositoryError", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedError := errors.New("error")

		mockRepo.On("GetAll").Return([]entities.Regency{}, expectedError)

		resultRegencies, err := regencyUseCase.GetAll()

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, resultRegencies)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedRegency := entities.Regency{ID: "1", Name: "Regency 1"}

		mockRepo.On("GetByID", "1").Return(expectedRegency, nil)

		resultRegency, err := regencyUseCase.GetByID("1")

		assert.NoError(t, err)
		assert.Equal(t, expectedRegency, resultRegency)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRepo := new(MockRegencyRepository)
		mockAPI := new(MockRegencyAPI)
		regencyUseCase := NewRegencyUsecase(mockRepo, mockAPI)

		expectedError := errors.New("error")

		mockRepo.On("GetByID", "1").Return(entities.Regency{}, expectedError)

		resultRegency, err := regencyUseCase.GetByID("1")

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Regency{}, resultRegency)
	})
}
