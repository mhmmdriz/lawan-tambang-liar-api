package report_upvote

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReportUpvoteRepository struct {
	mock.Mock
}

func (m *MockReportUpvoteRepository) ToggleUpvote(reportUpvote entities.ReportUpvote) (entities.ReportUpvote, string, error) {
	args := m.Called(reportUpvote)
	return args.Get(0).(entities.ReportUpvote), args.String(1), args.Error(2)
}

func TestToggleUpvote(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockReportUpvoteRepository := new(MockReportUpvoteRepository)

		reportUpvoteUseCase := NewReportUpvoteUseCase(mockReportUpvoteRepository)
		mockReportUpvoteRepository.On("ToggleUpvote", mock.Anything).Return(entities.ReportUpvote{}, "upvoted", nil)

		reportUpvote, status, err := reportUpvoteUseCase.ToggleUpvote(1, 1)

		assert.Nil(t, err)
		assert.Equal(t, "upvoted", status)
		assert.Equal(t, entities.ReportUpvote{}, reportUpvote)
	})

	t.Run("FailedAllFieldsMustBeFilled", func(t *testing.T) {
		mockReportUpvoteRepository := new(MockReportUpvoteRepository)

		reportUpvoteUseCase := NewReportUpvoteUseCase(mockReportUpvoteRepository)

		reportUpvote, status, err := reportUpvoteUseCase.ToggleUpvote(0, 0)

		assert.NotNil(t, err)
		assert.Equal(t, "", status)
		assert.Equal(t, entities.ReportUpvote{}, reportUpvote)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockReportUpvoteRepository := new(MockReportUpvoteRepository)

		reportUpvoteUseCase := NewReportUpvoteUseCase(mockReportUpvoteRepository)
		mockReportUpvoteRepository.On("ToggleUpvote", mock.Anything).Return(entities.ReportUpvote{}, "", constants.ErrInternalServerError)

		reportUpvote, status, err := reportUpvoteUseCase.ToggleUpvote(1, 1)

		assert.NotNil(t, err)
		assert.Equal(t, "", status)
		assert.Equal(t, entities.ReportUpvote{}, reportUpvote)
	})
}
