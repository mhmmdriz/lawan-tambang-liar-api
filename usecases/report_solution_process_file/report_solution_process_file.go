package report_solution_process_file

import (
	"lawan-tambang-liar/entities"
	"mime/multipart"
)

type ReportSolutionProcessFileUsecase struct {
	repository entities.ReportSolutionProcessFileRepositoryInterface
	gcs_api    entities.ReportSolutionProcessFileGCSAPIInterface
}

func NewReportSolutionProcessFileUsecase(repository entities.ReportSolutionProcessFileRepositoryInterface, gcs_api entities.ReportSolutionProcessFileGCSAPIInterface) *ReportSolutionProcessFileUsecase {
	return &ReportSolutionProcessFileUsecase{
		repository: repository,
		gcs_api:    gcs_api,
	}
}

func (u *ReportSolutionProcessFileUsecase) Create(files []*multipart.FileHeader, reportSolutionProcessID int) ([]entities.ReportSolutionProcessFile, error) {
	filepaths, err_upload := u.gcs_api.UploadFile(files)
	if err_upload != nil {
		return []entities.ReportSolutionProcessFile{}, err_upload
	}

	var reportFiles []*entities.ReportSolutionProcessFile
	for _, filepath := range filepaths {
		reportFile := &entities.ReportSolutionProcessFile{
			ReportSolutionProcessID: reportSolutionProcessID,
			Path:                    filepath,
		}
		reportFiles = append(reportFiles, reportFile)
	}

	err := u.repository.Create(reportFiles)

	if err != nil {
		return []entities.ReportSolutionProcessFile{}, err_upload
	}

	// Convert []*entities.ReportSolutionProcessFile to []entities.ReportSolutionProcessFile
	var convertedReportFiles []entities.ReportSolutionProcessFile
	for _, rf := range reportFiles {
		convertedReportFiles = append(convertedReportFiles, *rf)
	}

	return convertedReportFiles, nil
}

func (u *ReportSolutionProcessFileUsecase) Delete(reportSolutionProcessID int) ([]entities.ReportSolutionProcessFile, error) {
	reportFiles, err := u.repository.Delete(reportSolutionProcessID)

	if err != nil {
		return []entities.ReportSolutionProcessFile{}, err
	}

	var filePaths []string
	for _, file := range reportFiles {
		filePaths = append(filePaths, file.Path)
	}

	err_delete := u.gcs_api.DeleteFile(filePaths)
	if err_delete != nil {
		return []entities.ReportSolutionProcessFile{}, err_delete
	}

	return reportFiles, nil
}
