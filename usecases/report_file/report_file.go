package report_file

import (
	"lawan-tambang-liar/entities"
	"mime/multipart"
)

type ReportFileUseCase struct {
	repository entities.ReportFileRepositoryInterface
	gcs_api    entities.ReportFileGCSAPIInterface
}

func NewReportFileUseCase(repository entities.ReportFileRepositoryInterface, gcs_api entities.ReportFileGCSAPIInterface) *ReportFileUseCase {
	return &ReportFileUseCase{
		repository: repository,
		gcs_api:    gcs_api,
	}
}

func (u *ReportFileUseCase) Create(files []*multipart.FileHeader, report_id int) ([]entities.ReportFile, error) {
	filepaths, err_upload := u.gcs_api.UploadFile(files)
	if err_upload != nil {
		return []entities.ReportFile{}, err_upload
	}

	var reportFiles []*entities.ReportFile
	for _, filepath := range filepaths {
		reportFile := &entities.ReportFile{
			ReportID: report_id,
			Path:     filepath,
		}
		reportFiles = append(reportFiles, reportFile)
	}

	err := u.repository.Create(reportFiles)

	if err != nil {
		return []entities.ReportFile{}, err_upload
	}

	// Convert []*entities.ReportFile to []entities.ReportFile
	var convertedReportFiles []entities.ReportFile
	for _, rf := range reportFiles {
		convertedReportFiles = append(convertedReportFiles, *rf)
	}

	return convertedReportFiles, nil
}

func (u *ReportFileUseCase) Delete(report_id int) ([]entities.ReportFile, error) {
	reportFiles, err := u.repository.Delete(report_id)

	if err != nil {
		return []entities.ReportFile{}, err
	}

	err_delete := u.gcs_api.DeleteFile(reportFiles)
	if err_delete != nil {
		return []entities.ReportFile{}, err_delete
	}

	return reportFiles, nil
}
