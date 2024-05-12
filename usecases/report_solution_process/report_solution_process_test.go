package report_solution_process

import (
	"errors"
	"lawan-tambang-liar/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReportSolutionProcessRepository struct {
	mock.Mock
}

func (m *MockReportSolutionProcessRepository) Create(reportSolutionProcess *entities.ReportSolutionProcess) error {
	args := m.Called(reportSolutionProcess)
	return args.Error(0)
}

func (m *MockReportSolutionProcessRepository) GetByReportID(reportID int) ([]entities.ReportSolutionProcess, error) {
	args := m.Called(reportID)
	return args.Get(0).([]entities.ReportSolutionProcess), args.Error(1)
}

func (m *MockReportSolutionProcessRepository) Delete(reportID int, reportSolutionProcessStatus string) (entities.ReportSolutionProcess, error) {
	args := m.Called(reportID, reportSolutionProcessStatus)
	return args.Get(0).(entities.ReportSolutionProcess), args.Error(1)
}

func (m *MockReportSolutionProcessRepository) Update(reportSolutionProcess entities.ReportSolutionProcess) (entities.ReportSolutionProcess, error) {
	args := m.Called(reportSolutionProcess)
	return args.Get(0).(entities.ReportSolutionProcess), args.Error(1)
}

type MockAIReportSolutionAPI struct {
	mock.Mock
}

func (m *MockAIReportSolutionAPI) GetChatCompletion(messages []map[string]string) (string, error) {
	args := m.Called(messages)
	return args.String(0), args.Error(1)
}

func TestCreate(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("Create", mock.Anything).Return(nil)

		reportSolutionProcess := &entities.ReportSolutionProcess{
			ReportID: 1,
			AdminID:  1,
			Message:  "message",
			Status:   "status",
		}

		result, err := reportSolutionProcessUseCase.Create(reportSolutionProcess)

		assert.Nil(t, err)
		assert.Equal(t, reportSolutionProcess, &result)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)

		reportSolutionProcess := &entities.ReportSolutionProcess{
			ReportID: 0,
			AdminID:  0,
			Message:  "",
			Status:   "",
		}

		result, err := reportSolutionProcessUseCase.Create(reportSolutionProcess)

		assert.NotNil(t, err)
		assert.Equal(t, entities.ReportSolutionProcess{}, result)
	})

	t.Run("FailedCreateRepo", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("Create", mock.Anything).Return(errors.New("unexpected error"))

		reportSolutionProcess := &entities.ReportSolutionProcess{
			ReportID: 1,
			AdminID:  1,
			Message:  "message",
			Status:   "status",
		}

		result, err := reportSolutionProcessUseCase.Create(reportSolutionProcess)

		assert.NotNil(t, err)
		assert.Equal(t, entities.ReportSolutionProcess{}, result)
	})
}

func TestGetByReportID(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("GetByReportID", mock.Anything).Return([]entities.ReportSolutionProcess{}, nil)

		result, err := reportSolutionProcessUseCase.GetByReportID(1)

		assert.Nil(t, err)
		assert.Equal(t, []entities.ReportSolutionProcess{}, result)
	})

	t.Run("FailedRepo", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("GetByReportID", mock.Anything).Return([]entities.ReportSolutionProcess{}, errors.New("unexpected error"))

		result, err := reportSolutionProcessUseCase.GetByReportID(1)

		assert.NotNil(t, err)
		assert.Equal(t, []entities.ReportSolutionProcess{}, result)
	})
}

func TestDelete(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("Delete", mock.Anything, mock.Anything).Return(entities.ReportSolutionProcess{}, nil)

		result, err := reportSolutionProcessUseCase.Delete(1, "status")

		assert.Nil(t, err)
		assert.Equal(t, entities.ReportSolutionProcess{}, result)
	})

	t.Run("FailedRepo", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("Delete", mock.Anything, mock.Anything).Return(entities.ReportSolutionProcess{}, errors.New("unexpected error"))

		result, err := reportSolutionProcessUseCase.Delete(1, "status")

		assert.NotNil(t, err)
		assert.Equal(t, entities.ReportSolutionProcess{}, result)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("Update", mock.Anything).Return(entities.ReportSolutionProcess{}, nil)

		reportSolutionProcess := entities.ReportSolutionProcess{
			ReportID: 1,
			AdminID:  1,
			Message:  "message",
			Status:   "status",
		}

		result, err := reportSolutionProcessUseCase.Update(reportSolutionProcess)

		assert.Nil(t, err)
		assert.Equal(t, entities.ReportSolutionProcess{}, result)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)

		reportSolutionProcess := entities.ReportSolutionProcess{
			ReportID: 1,
			AdminID:  1,
			Message:  "",
			Status:   "",
		}

		result, err := reportSolutionProcessUseCase.Update(reportSolutionProcess)

		assert.NotNil(t, err)
		assert.Equal(t, entities.ReportSolutionProcess{}, result)
	})

	t.Run("FailedUpdateRepo", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockReportSolutionProcessRepository.On("Update", mock.Anything).Return(entities.ReportSolutionProcess{}, errors.New("unexpected error"))

		reportSolutionProcess := entities.ReportSolutionProcess{
			ReportID: 1,
			AdminID:  1,
			Message:  "message",
			Status:   "status",
		}

		result, err := reportSolutionProcessUseCase.Update(reportSolutionProcess)

		assert.NotNil(t, err)
		assert.Equal(t, entities.ReportSolutionProcess{}, result)
	})
}

func TestGetMessageRecommendation(t *testing.T) {

	t.Run("SuccessGetVerifyMessageRecommendation", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockAIReportSolutionAPI.On("GetChatCompletion", mock.Anything).Return("message", nil)

		result, err := reportSolutionProcessUseCase.GetMessageRecommendation("verify")

		assert.Nil(t, err)
		assert.Equal(t, "message", result)
	})

	t.Run("SuccessGetProgressMessageRecommendation", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockAIReportSolutionAPI.On("GetChatCompletion", mock.Anything).Return("message", nil)

		result, err := reportSolutionProcessUseCase.GetMessageRecommendation("progress")

		assert.Nil(t, err)
		assert.Equal(t, "message", result)
	})

	t.Run("SuccessGetFinishMessageRecommendation", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockAIReportSolutionAPI.On("GetChatCompletion", mock.Anything).Return("message", nil)

		result, err := reportSolutionProcessUseCase.GetMessageRecommendation("finish")

		assert.Nil(t, err)
		assert.Equal(t, "message", result)
	})

	t.Run("FailedRepo", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)
		mockAIReportSolutionAPI.On("GetChatCompletion", mock.Anything).Return("", errors.New("unexpected error"))

		result, err := reportSolutionProcessUseCase.GetMessageRecommendation("verify")

		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("FailedInvalidAction", func(t *testing.T) {
		mockReportSolutionProcessRepository := new(MockReportSolutionProcessRepository)
		mockAIReportSolutionAPI := new(MockAIReportSolutionAPI)

		reportSolutionProcessUseCase := NewReportSolutionProcessUseCase(mockReportSolutionProcessRepository, mockAIReportSolutionAPI)

		result, err := reportSolutionProcessUseCase.GetMessageRecommendation("invalid")

		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})
}
