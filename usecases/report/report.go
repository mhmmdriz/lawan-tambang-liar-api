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
