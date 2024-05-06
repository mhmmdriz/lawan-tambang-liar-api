package report_solution_process

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
)

type ReportSolutionProcessUseCase struct {
	repository entities.ReportSolutionProcessRepositoryInterface
}

func NewReportSolutionProcessUseCase(repository entities.ReportSolutionProcessRepositoryInterface) *ReportSolutionProcessUseCase {
	return &ReportSolutionProcessUseCase{
		repository: repository,
	}
}

func (u *ReportSolutionProcessUseCase) Create(reportSolutionProcess *entities.ReportSolutionProcess) (entities.ReportSolutionProcess, error) {
	if reportSolutionProcess.ReportID == 0 || reportSolutionProcess.AdminID == 0 || reportSolutionProcess.Message == "" || reportSolutionProcess.Status == "" {
		return entities.ReportSolutionProcess{}, constants.ErrAllFieldsMustBeFilled
	}

	err := u.repository.Create(reportSolutionProcess)

	if err != nil {
		return entities.ReportSolutionProcess{}, constants.ErrInternalServerError
	}

	return *reportSolutionProcess, nil
}

func (u *ReportSolutionProcessUseCase) GetByReportID(reportID int) ([]entities.ReportSolutionProcess, error) {
	reportSolutionProcesses, err := u.repository.GetByReportID(reportID)

	if err != nil {
		return nil, constants.ErrInternalServerError
	}

	return reportSolutionProcesses, nil
}

func (u *ReportSolutionProcessUseCase) Delete(reportID int, reportSolutionProcessStatus string) (entities.ReportSolutionProcess, error) {
	reportSolutionProcess, err := u.repository.Delete(reportID, reportSolutionProcessStatus)

	if err != nil {
		return entities.ReportSolutionProcess{}, err
	}

	return reportSolutionProcess, nil
}

func (u *ReportSolutionProcessUseCase) Update(reportSolutionProcess entities.ReportSolutionProcess) (entities.ReportSolutionProcess, error) {
	if reportSolutionProcess.Message == "" {
		return entities.ReportSolutionProcess{}, constants.ErrAllFieldsMustBeFilled
	}

	reportSolution, err := u.repository.Update(reportSolutionProcess)

	if err != nil {
		return entities.ReportSolutionProcess{}, err
	}

	return reportSolution, nil
}
