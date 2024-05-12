package district

import (
	"errors"
	"lawan-tambang-liar/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDistrictRepository struct {
	mock.Mock
}

func (m *MockDistrictRepository) AddDistrictsFromAPI(districts []entities.District) error {
	args := m.Called(districts)
	return args.Error(0)
}

func (m *MockDistrictRepository) GetAll(regencyID string) ([]entities.District, error) {
	args := m.Called(regencyID)
	return args.Get(0).([]entities.District), args.Error(1)
}

func (m *MockDistrictRepository) GetByID(id string) (entities.District, error) {
	args := m.Called(id)
	return args.Get(0).(entities.District), args.Error(1)
}

type MockDistrictAPI struct {
	mock.Mock
}

func (m *MockDistrictAPI) GetDistrictsDataFromAPI(regencyIDs []string) ([]entities.District, error) {
	args := m.Called(regencyIDs)
	return args.Get(0).([]entities.District), args.Error(1)
}

func TestSeedDistrictDBFromAPI(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		regencyIDs := []string{"regencyID1", "regencyID2"}

		expectedDistricts := []entities.District{
			{ID: "1", Name: "District 1"},
			{ID: "2", Name: "District 2"},
		}

		mockAPI.On("GetDistrictsDataFromAPI", regencyIDs).Return(expectedDistricts, nil)
		mockRepo.On("AddDistrictsFromAPI", expectedDistricts).Return(nil)

		resultDistricts, err := districtUseCase.SeedDistrictDBFromAPI(regencyIDs)

		assert.NoError(t, err)
		assert.Equal(t, expectedDistricts, resultDistricts)

		mockAPI.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedAPIError", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		regencyIDs := []string{"regencyID1", "regencyID2"}

		expectedError := errors.New("API error")

		mockAPI.On("GetDistrictsDataFromAPI", regencyIDs).Return([]entities.District{}, expectedError)

		resultDistricts, err := districtUseCase.SeedDistrictDBFromAPI(regencyIDs)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, resultDistricts)

		mockAPI.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "AddDistrictsFromAPI")
	})

	t.Run("FailedRepositoryError", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		regencyIDs := []string{"regencyID1", "regencyID2"}

		expectedDistricts := []entities.District{
			{ID: "1", Name: "District 1"},
			{ID: "2", Name: "District 2"},
		}

		expectedError := errors.New("Repository error")

		mockAPI.On("GetDistrictsDataFromAPI", regencyIDs).Return(expectedDistricts, nil)
		mockRepo.On("AddDistrictsFromAPI", expectedDistricts).Return(expectedError)

		resultDistricts, err := districtUseCase.SeedDistrictDBFromAPI(regencyIDs)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, resultDistricts)

		mockAPI.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		regencyID := "regencyID1"
		expectedDistricts := []entities.District{
			{ID: "1", Name: "District 1"},
			{ID: "2", Name: "District 2"},
		}

		mockRepo.On("GetAll", regencyID).Return(expectedDistricts, nil)

		resultDistricts, err := districtUseCase.GetAll(regencyID)

		assert.NoError(t, err)
		assert.Equal(t, expectedDistricts, resultDistricts)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedRepositoryError", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		regencyID := "regencyID1"
		expectedError := errors.New("Repository error")

		mockRepo.On("GetAll", regencyID).Return([]entities.District{}, expectedError)

		resultDistricts, err := districtUseCase.GetAll(regencyID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, resultDistricts)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		districtID := "1"
		expectedDistrict := entities.District{ID: "1", Name: "District 1"}

		mockRepo.On("GetByID", districtID).Return(expectedDistrict, nil)

		resultDistrict, err := districtUseCase.GetByID(districtID)

		assert.NoError(t, err)
		assert.Equal(t, expectedDistrict, resultDistrict)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedRepositoryError", func(t *testing.T) {
		mockRepo := new(MockDistrictRepository)
		mockAPI := new(MockDistrictAPI)
		districtUseCase := NewDistrictUseCase(mockRepo, mockAPI)

		districtID := "1"
		expectedError := errors.New("Repository error")

		mockRepo.On("GetByID", districtID).Return(entities.District{}, expectedError)

		resultDistrict, err := districtUseCase.GetByID(districtID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.District{}, resultDistrict)

		mockRepo.AssertExpectations(t)
	})
}
