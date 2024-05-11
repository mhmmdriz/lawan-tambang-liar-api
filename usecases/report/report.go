package report

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"strings"
)

type ReportUseCase struct {
	report_repository entities.ReportRepositoryInterface
	admin_repository  entities.AdminRepositoryInterface
	gmaps_api         entities.GoogleMapsAPIInterface
	ai_api            entities.AIReportSolutionAPIInterface
}

func NewReportUseCase(report_repository entities.ReportRepositoryInterface, admin_repository entities.AdminRepositoryInterface, gmaps_api entities.GoogleMapsAPIInterface, ai_api entities.AIReportAPIInterface) *ReportUseCase {
	return &ReportUseCase{
		report_repository: report_repository,
		admin_repository:  admin_repository,
		gmaps_api:         gmaps_api,
		ai_api:            ai_api,
	}
}

func (u *ReportUseCase) Create(report *entities.Report) (entities.Report, error) {
	if report.UserID == 0 || report.Title == "" || report.Description == "" || report.RegencyID == "" || report.DistrictID == "" || report.Address == "" {
		return entities.Report{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.report_repository.Create(report)

	if err != nil {
		if strings.HasSuffix(err.Error(), "REFERENCES `regencies` (`id`))") {
			return entities.Report{}, constants.ErrRegencyNotFound
		} else if strings.HasSuffix(err.Error(), "REFERENCES `districts` (`id`))") {
			return entities.Report{}, constants.ErrDistrictNotFound
		} else {
			return entities.Report{}, constants.ErrInternalServerError
		}
	}

	return *report, nil
}

func (u *ReportUseCase) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Report, error) {
	if limit == 0 || page == 0 {
		return nil, constants.ErrLimitAndPageMustBeFilled
	}

	reports, err := u.report_repository.GetPaginated(limit, page, search, filter, sortBy, sortType)

	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return reports, nil
}

func (u *ReportUseCase) GetByID(id int) (entities.Report, error) {
	report, err := u.report_repository.GetByID(id)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) Update(report entities.Report) (entities.Report, error) {
	if report.UserID == 0 || report.Title == "" || report.Description == "" || report.RegencyID == "" || report.DistrictID == "" || report.Address == "" {
		return entities.Report{}, constants.ErrAllFieldsMustBeFilled
	}

	report, err := u.report_repository.Update(report)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) Delete(reportID int, userID int) (entities.Report, error) {
	if reportID == 0 {
		return entities.Report{}, constants.ErrIDMustBeFilled
	}

	report, err := u.report_repository.Delete(reportID, userID)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) AdminDelete(reportID int) (entities.Report, error) {
	if reportID == 0 {
		return entities.Report{}, constants.ErrIDMustBeFilled
	}

	report, err := u.report_repository.AdminDelete(reportID)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) UpdateStatus(reportID int, status string) error {
	if reportID == 0 {
		return constants.ErrIDMustBeFilled
	}

	err := u.report_repository.UpdateStatus(reportID, status)

	if err != nil {
		return err
	}

	return nil
}

func (u *ReportUseCase) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	meta, err := u.report_repository.GetMetaData(limit, page, search, filter)

	if err != nil {
		return entities.Metadata{}, constants.ErrInternalServerError
	}

	return meta, nil
}

func (u *ReportUseCase) GetDistanceDuration(reportID int, adminID int) (entities.DistanceMatrix, error) {
	report, err := u.report_repository.GetByID(reportID)
	if err != nil {
		return entities.DistanceMatrix{}, err
	}

	admin, err := u.admin_repository.GetByID(adminID)
	if err != nil {
		return entities.DistanceMatrix{}, err
	}

	originAddress := admin.Address + ", " + admin.District.Name + ", " + admin.Regency.Name
	destinationAddress := report.Address + ", " + report.District.Name + ", " + report.Regency.Name

	distanceMatrix, err := u.gmaps_api.GetDistanceMatrix(originAddress, destinationAddress)

	if err != nil {
		return entities.DistanceMatrix{}, err
	}

	return distanceMatrix, nil
}

func (u *ReportUseCase) GetDescriptionRecommendation(location string) (string, error) {
	if location == "" {
		return "", constants.ErrAllFieldsMustBeFilled
	}

	messages := []map[string]string{
		{"role": "assistant", "content": "Anda sebagai masyarakat pengguna website lawan tambang liar dan bertugas untuk membuat laporan tambang liar"},
		{"role": "user", "content": "Saya masyarakat pengguna website lawan tambang liar di Provinsi Kepulauan Bangka Belitung. Berikan saya contoh deskripsi yang baik saat membuat laporan tambang liar ! Dan saya akan mengirimkan bukti fotonya juga. Tambang liar tersebut berada di" + location},
	}

	content, err := u.ai_api.GetChatCompletion(messages)

	if err != nil {
		return "", constants.ErrInternalServerError
	}

	return content, nil
}

func (u *ReportUseCase) IncreaseUpvote(reportID int) error {
	if reportID == 0 {
		return constants.ErrIDMustBeFilled
	}

	err := u.report_repository.IncreaseUpvote(reportID)

	if err != nil {
		return err
	}

	return nil
}

func (u *ReportUseCase) DecreaseUpvote(reportID int) error {
	if reportID == 0 {
		return constants.ErrIDMustBeFilled
	}

	err := u.report_repository.DecreaseUpvote(reportID)

	if err != nil {
		return err
	}

	return nil
}
