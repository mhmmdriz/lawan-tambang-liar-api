package report

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
)

type ReportUseCase struct {
	repository entities.ReportRepositoryInterface
}

func NewReportUseCase(repository entities.ReportRepositoryInterface) *ReportUseCase {
	return &ReportUseCase{
		repository: repository,
	}
}

func (u *ReportUseCase) Create(report *entities.Report) (entities.Report, error) {
	if report.UserID == 0 || report.Title == "" || report.Description == "" || report.RegencyID == "" || report.DistrictID == "" || report.Address == "" {
		return entities.Report{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Create(report)

	if err != nil {
		return entities.Report{}, constants.ErrInternalServerError
	}

	return *report, nil
}

func (u *ReportUseCase) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sort_by string, sort_type string) ([]entities.Report, error) {
	if limit == 0 || page == 0 {
		return nil, constants.ErrLimitAndPageMustBeFilled
	}

	reports, err := u.repository.GetPaginated(limit, page, search, filter, sort_by, sort_type)

	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return reports, nil
}

func (u *ReportUseCase) GetByID(id int) (entities.Report, error) {
	report, err := u.repository.GetByID(id)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) Update(report entities.Report) (entities.Report, error) {
	if report.ID == 0 {
		return entities.Report{}, constants.ErrIDMustBeFilled
	}

	report, err := u.repository.Update(report)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) Delete(report_id int, user_id int) (entities.Report, error) {
	if report_id == 0 {
		return entities.Report{}, constants.ErrIDMustBeFilled
	}

	report, err := u.repository.Delete(report_id, user_id)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) AdminDelete(report_id int) (entities.Report, error) {
	if report_id == 0 {
		return entities.Report{}, constants.ErrIDMustBeFilled
	}

	report, err := u.repository.AdminDelete(report_id)

	if err != nil {
		return entities.Report{}, err
	}

	return report, nil
}

func (u *ReportUseCase) UpdateStatus(report_id int, status string) error {
	if report_id == 0 {
		return constants.ErrIDMustBeFilled
	}

	err := u.repository.UpdateStatus(report_id, status)

	if err != nil {
		return err
	}

	return nil
}

func (u *ReportUseCase) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	meta, err := u.repository.GetMetaData(limit, page, search, filter)

	if err != nil {
		return entities.Metadata{}, constants.ErrInternalServerError
	}

	return meta, nil
}
