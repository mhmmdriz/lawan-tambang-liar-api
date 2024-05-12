package report_file_test

import (
	"errors"
	"fmt"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/usecases/report_file"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReportFileRepository struct {
	mock.Mock
}

func (m *MockReportFileRepository) Create(reportFiles []*entities.ReportFile) error {
	args := m.Called(reportFiles)
	return args.Error(0)
}

func (m *MockReportFileRepository) Delete(report_id int) ([]entities.ReportFile, error) {
	args := m.Called(report_id)
	return args.Get(0).([]entities.ReportFile), args.Error(1)
}

type MockReportFileGCSAPI struct {
	mock.Mock
}

func (m *MockReportFileGCSAPI) UploadFile(files []*multipart.FileHeader) ([]string, error) {
	args := m.Called(files)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockReportFileGCSAPI) DeleteFile(filepaths []string) error {
	args := m.Called(filepaths)
	return args.Error(0)
}

func TestCreate(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportFileRepository := new(MockReportFileRepository)
		mockReportFileGCSAPI := new(MockReportFileGCSAPI)

		reportFileUseCase := report_file.NewReportFileUseCase(mockReportFileRepository, mockReportFileGCSAPI)
		mockReportFileRepository.On("Create", mock.Anything).Return(nil)
		mockReportFileGCSAPI.On("UploadFile", mock.Anything).Return([]string{"file1", "file2"}, nil)

		files := []*multipart.FileHeader{
			{
				Filename: "file1",
			},
			{
				Filename: "file2",
			},
		}

		reportFiles, err := reportFileUseCase.Create(files, 1)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(reportFiles))
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockReportFileRepository := new(MockReportFileRepository)
		mockReportFileGCSAPI := new(MockReportFileGCSAPI)

		files := []*multipart.FileHeader{
			{
				Filename: "file1",
			},
			{
				Filename: "file2",
			},
		}

		reportFileUseCase := report_file.NewReportFileUseCase(mockReportFileRepository, mockReportFileGCSAPI)
		mockReportFileGCSAPI.On("UploadFile", mock.Anything).Return([]string{"file1", "file2"}, nil)
		mockReportFileRepository.On("Create", mock.Anything).Return(errors.New("failed to create report file"))

		reportFiles, err := reportFileUseCase.Create(files, 1)

		fmt.Println(err)
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(reportFiles))
	})

	t.Run("FailedReportFileGCSAPIError", func(t *testing.T) {
		mockReportFileRepository := new(MockReportFileRepository)
		mockReportFileGCSAPI := new(MockReportFileGCSAPI)

		reportFileUseCase := report_file.NewReportFileUseCase(mockReportFileRepository, mockReportFileGCSAPI)
		mockReportFileGCSAPI.On("UploadFile", mock.Anything).Return([]string{}, errors.New("failed to upload file"))

		files := []*multipart.FileHeader{
			{
				Filename: "file1",
			},
		}

		reportFiles, err := reportFileUseCase.Create(files, 1)

		assert.NotNil(t, err)
		assert.Equal(t, 0, len(reportFiles))
	})
}

func TestDelete(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		mockReportFileRepository := new(MockReportFileRepository)
		mockReportFileGCSAPI := new(MockReportFileGCSAPI)

		reportFileUseCase := report_file.NewReportFileUseCase(mockReportFileRepository, mockReportFileGCSAPI)
		mockReportFileRepository.On("Delete", mock.Anything).Return([]entities.ReportFile{
			{
				ReportID: 1,
				Path:     "file1",
			},
			{
				ReportID: 1,
				Path:     "file2",
			},
		}, nil)
		mockReportFileGCSAPI.On("DeleteFile", mock.Anything).Return(nil)

		reportFiles, err := reportFileUseCase.Delete(1)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(reportFiles))
	})

	t.Run("FailedRepoError", func(t *testing.T) {
		mockReportFileRepository := new(MockReportFileRepository)
		mockReportFileGCSAPI := new(MockReportFileGCSAPI)

		reportFileUseCase := report_file.NewReportFileUseCase(mockReportFileRepository, mockReportFileGCSAPI)
		mockReportFileRepository.On("Delete", mock.Anything).Return([]entities.ReportFile{}, errors.New("failed to delete report file"))

		reportFiles, err := reportFileUseCase.Delete(1)

		assert.NotNil(t, err)
		assert.Equal(t, 0, len(reportFiles))
	})

	t.Run("FailedReportFileGCSAPIError", func(t *testing.T) {
		mockReportFileRepository := new(MockReportFileRepository)
		mockReportFileGCSAPI := new(MockReportFileGCSAPI)

		reportFileUseCase := report_file.NewReportFileUseCase(mockReportFileRepository, mockReportFileGCSAPI)
		mockReportFileRepository.On("Delete", mock.Anything).Return([]entities.ReportFile{
			{
				ReportID: 1,
				Path:     "file1",
			},
		}, nil)
		mockReportFileGCSAPI.On("DeleteFile", mock.Anything).Return(errors.New("failed to delete file"))

		reportFiles, err := reportFileUseCase.Delete(1)

		assert.NotNil(t, err)
		assert.Equal(t, 0, len(reportFiles))
	})
}
