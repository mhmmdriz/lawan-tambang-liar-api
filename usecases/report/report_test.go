package report

import (
	"errors"
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReportRepository struct {
	mock.Mock
}

func (m *MockReportRepository) Create(report *entities.Report) error {
	args := m.Called(report)
	return args.Error(0)
}

func (m *MockReportRepository) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Report, error) {
	args := m.Called(limit, page, search, filter, sortBy, sortType)
	return args.Get(0).([]entities.Report), args.Error(1)
}

func (m *MockReportRepository) GetByID(id int) (entities.Report, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Report), args.Error(1)
}

func (m *MockReportRepository) Update(report entities.Report) (entities.Report, error) {
	args := m.Called(report)
	return args.Get(0).(entities.Report), args.Error(1)
}

func (m *MockReportRepository) Delete(reportID int, userID int) (entities.Report, error) {
	args := m.Called(reportID, userID)
	return args.Get(0).(entities.Report), args.Error(1)
}

func (m *MockReportRepository) AdminDelete(id int) (entities.Report, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Report), args.Error(1)
}

func (m *MockReportRepository) UpdateStatus(id int, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockReportRepository) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	args := m.Called(limit, page, search, filter)
	return args.Get(0).(entities.Metadata), args.Error(1)
}

func (m *MockReportRepository) IncreaseUpvote(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockReportRepository) DecreaseUpvote(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) GetByID(id int) (entities.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Admin), args.Error(1)
}

type MockAIReportAPI struct {
	mock.Mock
}

func (m *MockAIReportAPI) GetChatCompletion(messages []map[string]string) (string, error) {
	args := m.Called(messages)
	return args.String(0), args.Error(1)
}

type MockGMapsAPI struct {
	mock.Mock
}

func (m *MockGMapsAPI) GetDistanceMatrix(origins string, destinations string) (entities.DistanceMatrix, error) {
	args := m.Called(origins, destinations)
	return args.Get(0).(entities.DistanceMatrix), args.Error(1)
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockReportRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockReportRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		mockReportRepo.On("Create", &report).Return(nil)

		resultReport, err := reportUseCase.Create(&report)

		assert.NoError(t, err)
		assert.Equal(t, report, resultReport)

		mockReportRepo.AssertExpectations(t)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockReportRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockReportRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			UserID:      1,
			Title:       "",
			Description: "",
			RegencyID:   "",
			DistrictID:  "",
			Address:     "",
		}

		resultReport, err := reportUseCase.Create(&report)

		assert.Error(t, err)
		assert.Equal(t, entities.Report{}, resultReport)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		expectedError := constants.ErrInternalServerError
		mockRepo.On("Create", &report).Return(expectedError)

		resultReport, err := reportUseCase.Create(&report)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedRegencyNotFound", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		expectedError := errors.New("Error REFERENCES `regencies` (`id`))")
		mockRepo.On("Create", &report).Return(expectedError)

		resultReport, err := reportUseCase.Create(&report)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrRegencyNotFound, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedDistrictNotFound", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		expectedError := errors.New("Error REFERENCES `districts` (`id`))")
		mockRepo.On("Create", &report).Return(expectedError)

		resultReport, err := reportUseCase.Create(&report)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrDistrictNotFound, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

}

func TestGetPaginated(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		limit := 10
		page := 1
		search := ""
		filter := map[string]interface{}{}
		sortBy := ""
		sortType := ""

		expectedReports := []entities.Report{
			{ID: 1, Title: "Title 1"},
			{ID: 2, Title: "Title 2"},
		}

		mockRepo.On("GetPaginated", limit, page, search, filter, sortBy, sortType).Return(expectedReports, nil)

		resultReports, err := reportUseCase.GetPaginated(limit, page, search, filter, sortBy, sortType)

		assert.NoError(t, err)
		assert.Equal(t, expectedReports, resultReports)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedLimitAndPageMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		limit := 0
		page := 0
		search := ""
		filter := map[string]interface{}{}
		sortBy := ""
		sortType := ""

		resultReports, err := reportUseCase.GetPaginated(limit, page, search, filter, sortBy, sortType)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrLimitAndPageMustBeFilled, err)
		assert.Nil(t, resultReports)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		limit := 10
		page := 1
		search := ""
		filter := map[string]interface{}{}
		sortBy := ""
		sortType := ""

		expectedError := constants.ErrInternalServerError
		mockRepo.On("GetPaginated", limit, page, search, filter, sortBy, sortType).Return([]entities.Report{}, expectedError)

		resultReports, err := reportUseCase.GetPaginated(limit, page, search, filter, sortBy, sortType)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, resultReports)

		mockRepo.AssertExpectations(t)
	})

}

func TestGetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1
		expectedReport := entities.Report{ID: 1, Title: "Title"}

		mockRepo.On("GetByID", id).Return(expectedReport, nil)

		resultReport, err := reportUseCase.GetByID(id)

		assert.NoError(t, err)
		assert.Equal(t, expectedReport, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1
		expectedError := constants.ErrInternalServerError

		mockRepo.On("GetByID", id).Return(entities.Report{}, expectedError)

		resultReport, err := reportUseCase.GetByID(id)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			ID:          1,
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		mockRepo.On("Update", report).Return(report, nil)

		resultReport, err := reportUseCase.Update(report)

		assert.NoError(t, err)
		assert.Equal(t, report, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			ID:          1,
			UserID:      1,
			Title:       "",
			Description: "",
			RegencyID:   "",
			DistrictID:  "",
			Address:     "",
		}

		resultReport, err := reportUseCase.Update(report)

		assert.Error(t, err)
		assert.Equal(t, entities.Report{}, resultReport)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			ID:          1,
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		expectedError := constants.ErrInternalServerError
		mockRepo.On("Update", report).Return(entities.Report{}, expectedError)

		resultReport, err := reportUseCase.Update(report)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedReportNotFound", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			ID:          1,
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		expectedError := constants.ErrReportNotFound
		mockRepo.On("Update", report).Return(entities.Report{}, expectedError)

		resultReport, err := reportUseCase.Update(report)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedUnauthorized", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		report := entities.Report{
			ID:          1,
			UserID:      1,
			Title:       "Title",
			Description: "Description",
			RegencyID:   "regencyID",
			DistrictID:  "districtID",
			Address:     "Address",
		}

		mockRepo.On("Update", report).Return(entities.Report{}, constants.ErrUnauthorized)

		resultReport, err := reportUseCase.Update(report)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUnauthorized, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		userID := 1
		expectedReport := entities.Report{ID: 1, Title: "Title"}

		mockRepo.On("Delete", reportID, userID).Return(expectedReport, nil)

		resultReport, err := reportUseCase.Delete(reportID, userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedReport, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedIDMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 0
		userID := 1

		resultReport, err := reportUseCase.Delete(reportID, userID)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrIDMustBeFilled, err)
		assert.Equal(t, entities.Report{}, resultReport)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		userID := 1

		expectedError := constants.ErrInternalServerError
		mockRepo.On("Delete", reportID, userID).Return(entities.Report{}, expectedError)

		resultReport, err := reportUseCase.Delete(reportID, userID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedReportNotFound", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		userID := 1

		expectedError := constants.ErrReportNotFound
		mockRepo.On("Delete", reportID, userID).Return(entities.Report{}, expectedError)

		resultReport, err := reportUseCase.Delete(reportID, userID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedUnauthorized", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		userID := 1

		mockRepo.On("Delete", reportID, userID).Return(entities.Report{}, constants.ErrUnauthorized)

		resultReport, err := reportUseCase.Delete(reportID, userID)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUnauthorized, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})
}

func TestAdminDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		expectedReport := entities.Report{ID: 1, Title: "Title"}

		mockRepo.On("AdminDelete", reportID).Return(expectedReport, nil)

		resultReport, err := reportUseCase.AdminDelete(reportID)

		assert.NoError(t, err)
		assert.Equal(t, expectedReport, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedIDMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 0

		resultReport, err := reportUseCase.AdminDelete(reportID)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrIDMustBeFilled, err)
		assert.Equal(t, entities.Report{}, resultReport)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1

		expectedError := constants.ErrInternalServerError
		mockRepo.On("AdminDelete", reportID).Return(entities.Report{}, expectedError)

		resultReport, err := reportUseCase.AdminDelete(reportID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedReportNotFound", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1

		expectedError := constants.ErrReportNotFound
		mockRepo.On("AdminDelete", reportID).Return(entities.Report{}, expectedError)

		resultReport, err := reportUseCase.AdminDelete(reportID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Report{}, resultReport)

		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1
		status := "status"

		mockRepo.On("UpdateStatus", id, status).Return(nil)

		err := reportUseCase.UpdateStatus(id, status)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedIDMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 0
		status := "status"

		err := reportUseCase.UpdateStatus(id, status)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrIDMustBeFilled, err)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1
		status := "status"

		expectedError := constants.ErrInternalServerError
		mockRepo.On("UpdateStatus", id, status).Return(expectedError)

		err := reportUseCase.UpdateStatus(id, status)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetMetaData(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		limit := 10
		page := 1
		search := ""
		filter := map[string]interface{}{}

		expectedMetadata := entities.Metadata{
			TotalData: 10,
			Pagination: entities.Pagination{
				FirstPage:        1,
				LastPage:         1,
				CurrentPage:      1,
				TotalDataPerPage: 10,
				PrevPage:         0,
				NextPage:         0,
			},
		}

		mockRepo.On("GetMetaData", limit, page, search, filter).Return(expectedMetadata, nil)

		resultMetadata, err := reportUseCase.GetMetaData(limit, page, search, filter)

		assert.NoError(t, err)
		assert.Equal(t, expectedMetadata, resultMetadata)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		limit := 10
		page := 1
		search := ""
		filter := map[string]interface{}{}

		expectedError := constants.ErrInternalServerError
		mockRepo.On("GetMetaData", limit, page, search, filter).Return(entities.Metadata{}, expectedError)

		resultMetadata, err := reportUseCase.GetMetaData(limit, page, search, filter)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.Metadata{}, resultMetadata)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetDistanceDuration(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockGMaps := new(MockGMapsAPI)
		mockAI := new(MockAIReportAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		adminID := 1
		origin := "Address Origin, District, Regency"
		destination := "Address Destination, District, Regency"

		expectedDistanceMatrix := entities.DistanceMatrix{
			DestinationAddress: "destination",
			OriginAddress:      "origin",
			Distance:           "1 km",
			Duration:           "1 hour",
		}

		expectedAdmin := entities.Admin{
			ID: adminID,
			Regency: entities.Regency{
				Name: "Regency",
			},
			District: entities.District{
				Name: "District",
			},
			Address: "Address Origin",
		}

		expectedReport := entities.Report{
			ID: reportID,
			Regency: entities.Regency{
				Name: "Regency",
			},
			District: entities.District{
				Name: "District",
			},
			Address: "Address Destination",
		}

		mockAdminRepo.On("GetByID", adminID).Return(expectedAdmin, nil)
		mockRepo.On("GetByID", reportID).Return(expectedReport, nil)
		mockGMaps.On("GetDistanceMatrix", origin, destination).Return(expectedDistanceMatrix, nil)

		resultDistanceMatrix, err := reportUseCase.GetDistanceDuration(reportID, adminID)

		assert.NoError(t, err)
		assert.Equal(t, expectedDistanceMatrix, resultDistanceMatrix)

		mockGMaps.AssertExpectations(t)
	})

	t.Run("FailedAdminRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockGMaps := new(MockGMapsAPI)
		mockAI := new(MockAIReportAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		adminID := 1

		expectedReport := entities.Report{
			ID: reportID,
			Regency: entities.Regency{
				Name: "Regency",
			},
			District: entities.District{
				Name: "District",
			},
			Address: "Address Destination",
		}

		expectedError := constants.ErrInternalServerError
		mockAdminRepo.On("GetByID", adminID).Return(entities.Admin{}, expectedError)
		mockRepo.On("GetByID", reportID).Return(expectedReport, nil)

		resultDistanceMatrix, err := reportUseCase.GetDistanceDuration(reportID, adminID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.DistanceMatrix{}, resultDistanceMatrix)

		mockAdminRepo.AssertExpectations(t)
	})

	t.Run("FailedReportRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockGMaps := new(MockGMapsAPI)
		mockAI := new(MockAIReportAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		adminID := 1

		expectedError := constants.ErrInternalServerError
		mockRepo.On("GetByID", reportID).Return(entities.Report{}, expectedError)

		resultDistanceMatrix, err := reportUseCase.GetDistanceDuration(reportID, adminID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.DistanceMatrix{}, resultDistanceMatrix)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedGMapsError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockGMaps := new(MockGMapsAPI)
		mockAI := new(MockAIReportAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		reportID := 1
		adminID := 1
		origin := "Address Origin, District, Regency"
		destination := "Address Destination, District, Regency"

		expectedAdmin := entities.Admin{
			ID: adminID,
			Regency: entities.Regency{
				Name: "Regency",
			},
			District: entities.District{
				Name: "District",
			},
			Address: "Address Origin",
		}

		expectedReport := entities.Report{
			ID: reportID,
			Regency: entities.Regency{
				Name: "Regency",
			},
			District: entities.District{
				Name: "District",
			},
			Address: "Address Destination",
		}

		expectedError := constants.ErrInternalServerError
		mockAdminRepo.On("GetByID", adminID).Return(expectedAdmin, nil)
		mockRepo.On("GetByID", reportID).Return(expectedReport, nil)
		mockGMaps.On("GetDistanceMatrix", origin, destination).Return(entities.DistanceMatrix{}, expectedError)

		resultDistanceMatrix, err := reportUseCase.GetDistanceDuration(reportID, adminID)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, entities.DistanceMatrix{}, resultDistanceMatrix)

		mockGMaps.AssertExpectations(t)
	})
}

func TestGetDescriptionRecommendation(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockGMaps := new(MockGMapsAPI)
		mockAI := new(MockAIReportAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		location := "location"

		messages := []map[string]string{
			{"role": "assistant", "content": "Anda sebagai masyarakat pengguna website lawan tambang liar dan bertugas untuk membuat laporan tambang liar"},
			{"role": "user", "content": "Saya masyarakat pengguna website lawan tambang liar di Provinsi Kepulauan Bangka Belitung. Berikan saya contoh deskripsi yang baik saat membuat laporan tambang liar ! Dan saya akan mengirimkan bukti fotonya juga. Tambang liar tersebut berada di" + location},
		}

		expectedDescription := "description"

		mockAI.On("GetChatCompletion", messages).Return(expectedDescription, nil)

		resultDescription, err := reportUseCase.GetDescriptionRecommendation(location)

		assert.NoError(t, err)
		assert.Equal(t, expectedDescription, resultDescription)

		mockAI.AssertExpectations(t)
	})

	t.Run("FailedLocationMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockGMaps := new(MockGMapsAPI)
		mockAI := new(MockAIReportAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		location := ""

		resultDescription, err := reportUseCase.GetDescriptionRecommendation(location)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrAllFieldsMustBeFilled, err)
		assert.Equal(t, "", resultDescription)
	})

	t.Run("FailedAIError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockGMaps := new(MockGMapsAPI)
		mockAI := new(MockAIReportAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		location := "location"

		messages := []map[string]string{
			{"role": "assistant", "content": "Anda sebagai masyarakat pengguna website lawan tambang liar dan bertugas untuk membuat laporan tambang liar"},
			{"role": "user", "content": "Saya masyarakat pengguna website lawan tambang liar di Provinsi Kepulauan Bangka Belitung. Berikan saya contoh deskripsi yang baik saat membuat laporan tambang liar ! Dan saya akan mengirimkan bukti fotonya juga. Tambang liar tersebut berada di" + location},
		}

		expectedError := constants.ErrInternalServerError

		mockAI.On("GetChatCompletion", messages).Return("", expectedError)

		resultDescription, err := reportUseCase.GetDescriptionRecommendation(location)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, "", resultDescription)

		mockAI.AssertExpectations(t)
	})
}

func TestIncreaseUpvote(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1

		mockRepo.On("IncreaseUpvote", id).Return(nil)

		err := reportUseCase.IncreaseUpvote(id)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedIDMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 0

		err := reportUseCase.IncreaseUpvote(id)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrIDMustBeFilled, err)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1

		expectedError := constants.ErrInternalServerError
		mockRepo.On("IncreaseUpvote", id).Return(expectedError)

		err := reportUseCase.IncreaseUpvote(id)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)

		mockRepo.AssertExpectations(t)
	})
}

func TestDecreaseUpvote(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1

		mockRepo.On("DecreaseUpvote", id).Return(nil)

		err := reportUseCase.DecreaseUpvote(id)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FailedIDMustBeFilled", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 0

		err := reportUseCase.DecreaseUpvote(id)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrIDMustBeFilled, err)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockRepo := new(MockReportRepository)
		mockAdminRepo := new(MockAdminRepository)
		mockAI := new(MockAIReportAPI)
		mockGMaps := new(MockGMapsAPI)
		reportUseCase := NewReportUseCase(mockRepo, mockAdminRepo, mockGMaps, mockAI)

		id := 1

		expectedError := constants.ErrInternalServerError
		mockRepo.On("DecreaseUpvote", id).Return(expectedError)

		err := reportUseCase.DecreaseUpvote(id)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)

		mockRepo.AssertExpectations(t)
	})
}
