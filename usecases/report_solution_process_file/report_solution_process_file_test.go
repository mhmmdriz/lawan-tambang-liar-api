package report_solution_process_file

import (
	"errors"
	"lawan-tambang-liar/entities"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReportSolutionProcessFileRepository struct {
	mock.Mock
}

func (m *MockReportSolutionProcessFileRepository) Create(reportFiles []*entities.ReportSolutionProcessFile) error {
	args := m.Called(reportFiles)
	return args.Error(0)
}

func (m *MockReportSolutionProcessFileRepository) Delete(reportSolutionProcessID int) ([]entities.ReportSolutionProcessFile, error) {
	args := m.Called(reportSolutionProcessID)
	return args.Get(0).([]entities.ReportSolutionProcessFile), args.Error(1)
}

type MockReportSolutionProcessFileGCSAPI struct {
	mock.Mock
}

func (m *MockReportSolutionProcessFileGCSAPI) UploadFile(files []*multipart.FileHeader) ([]string, error) {
	args := m.Called(files)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockReportSolutionProcessFileGCSAPI) DeleteFile(filepaths []string) error {
	args := m.Called(filepaths)
	return args.Error(0)
}

func TestCreate(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportSolutionProcessFileRepository := new(MockReportSolutionProcessFileRepository)
		mockReportSolutionProcessFileGCSAPI := new(MockReportSolutionProcessFileGCSAPI)

		reportSolutionProcessFileUsecase := NewReportSolutionProcessFileUsecase(mockReportSolutionProcessFileRepository, mockReportSolutionProcessFileGCSAPI)
		mockReportSolutionProcessFileRepository.On("Create", mock.Anything).Return(nil)
		mockReportSolutionProcessFileGCSAPI.On("UploadFile", mock.Anything).Return([]string{"file1", "file2"}, nil)

		files := []*multipart.FileHeader{
			{
				Filename: "file1",
			},
			{
				Filename: "file2",
			},
		}

		_, err := reportSolutionProcessFileUsecase.Create(files, 1)

		assert.Nil(t, err)
		mockReportSolutionProcessFileRepository.AssertExpectations(t)
		mockReportSolutionProcessFileGCSAPI.AssertExpectations(t)
	})

	t.Run("FailedGCSAPIError", func(t *testing.T) {
		mockReportSolutionProcessFileRepository := new(MockReportSolutionProcessFileRepository)
		mockReportSolutionProcessFileGCSAPI := new(MockReportSolutionProcessFileGCSAPI)

		reportSolutionProcessFileUsecase := NewReportSolutionProcessFileUsecase(mockReportSolutionProcessFileRepository, mockReportSolutionProcessFileGCSAPI)
		mockReportSolutionProcessFileGCSAPI.On("UploadFile", mock.Anything).Return([]string{}, errors.New("error"))

		files := []*multipart.FileHeader{
			{
				Filename: "file1",
			},
			{
				Filename: "file2",
			},
		}

		_, err := reportSolutionProcessFileUsecase.Create(files, 1)

		assert.NotNil(t, err)
		mockReportSolutionProcessFileRepository.AssertExpectations(t)
		mockReportSolutionProcessFileGCSAPI.AssertExpectations(t)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockReportSolutionProcessFileRepository := new(MockReportSolutionProcessFileRepository)
		mockReportSolutionProcessFileGCSAPI := new(MockReportSolutionProcessFileGCSAPI)

		reportSolutionProcessFileUsecase := NewReportSolutionProcessFileUsecase(mockReportSolutionProcessFileRepository, mockReportSolutionProcessFileGCSAPI)
		mockReportSolutionProcessFileGCSAPI.On("UploadFile", mock.Anything).Return([]string{"file1", "file2"}, nil)
		mockReportSolutionProcessFileRepository.On("Create", mock.Anything).Return(errors.New("error"))

		files := []*multipart.FileHeader{
			{
				Filename: "file1",
			},
			{
				Filename: "file2",
			},
		}

		_, err := reportSolutionProcessFileUsecase.Create(files, 1)

		assert.NotNil(t, err)
		mockReportSolutionProcessFileRepository.AssertExpectations(t)
		mockReportSolutionProcessFileGCSAPI.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportSolutionProcessFileRepository := new(MockReportSolutionProcessFileRepository)
		mockReportSolutionProcessFileGCSAPI := new(MockReportSolutionProcessFileGCSAPI)

		reportSolutionProcessFileUsecase := NewReportSolutionProcessFileUsecase(mockReportSolutionProcessFileRepository, mockReportSolutionProcessFileGCSAPI)
		mockReportSolutionProcessFileRepository.On("Delete", mock.Anything).Return([]entities.ReportSolutionProcessFile{}, nil)
		mockReportSolutionProcessFileGCSAPI.On("DeleteFile", mock.Anything).Return(nil)

		_, err := reportSolutionProcessFileUsecase.Delete(1)

		assert.Nil(t, err)
		mockReportSolutionProcessFileRepository.AssertExpectations(t)
		mockReportSolutionProcessFileGCSAPI.AssertExpectations(t)
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockReportSolutionProcessFileRepository := new(MockReportSolutionProcessFileRepository)
		mockReportSolutionProcessFileGCSAPI := new(MockReportSolutionProcessFileGCSAPI)

		reportSolutionProcessFileUsecase := NewReportSolutionProcessFileUsecase(mockReportSolutionProcessFileRepository, mockReportSolutionProcessFileGCSAPI)
		mockReportSolutionProcessFileRepository.On("Delete", mock.Anything).Return([]entities.ReportSolutionProcessFile{}, errors.New("error"))

		_, err := reportSolutionProcessFileUsecase.Delete(1)

		assert.NotNil(t, err)
		mockReportSolutionProcessFileRepository.AssertExpectations(t)
		mockReportSolutionProcessFileGCSAPI.AssertExpectations(t)
	})

	t.Run("FailedGCSAPIError", func(t *testing.T) {
		mockReportSolutionProcessFileRepository := new(MockReportSolutionProcessFileRepository)
		mockReportSolutionProcessFileGCSAPI := new(MockReportSolutionProcessFileGCSAPI)

		reportSolutionProcessFileUsecase := NewReportSolutionProcessFileUsecase(mockReportSolutionProcessFileRepository, mockReportSolutionProcessFileGCSAPI)
		mockReportSolutionProcessFileRepository.On("Delete", mock.Anything).Return([]entities.ReportSolutionProcessFile{
			{
				ReportSolutionProcessID: 1,
				Path:                    "file1",
			},
			{
				ReportSolutionProcessID: 1,
				Path:                    "file2",
			},
		}, nil)
		mockReportSolutionProcessFileGCSAPI.On("DeleteFile", mock.Anything).Return(errors.New("error"))

		_, err := reportSolutionProcessFileUsecase.Delete(1)

		assert.NotNil(t, err)
		mockReportSolutionProcessFileRepository.AssertExpectations(t)
		mockReportSolutionProcessFileGCSAPI.AssertExpectations(t)
	})
}
