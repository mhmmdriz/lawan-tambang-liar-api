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
