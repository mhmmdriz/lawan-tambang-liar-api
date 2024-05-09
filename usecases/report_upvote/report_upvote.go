package report_upvote

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
)

type ReportUpvoteUseCase struct {
	repository entities.ReportUpvoteRepositoryInterface
}

func NewReportUpvoteUseCase(repository entities.ReportUpvoteRepositoryInterface) *ReportUpvoteUseCase {
	return &ReportUpvoteUseCase{
		repository: repository,
	}
}

func (u *ReportUpvoteUseCase) ToggleUpvote(userID int, reportID int) (entities.ReportUpvote, string, error) {
	if userID == 0 || reportID == 0 {
		return entities.ReportUpvote{}, "", constants.ErrIDMustBeFilled
	}

	reportUpvote := entities.ReportUpvote{
		UserID:   userID,
		ReportID: reportID,
	}

	reportUpvote, status, err := u.repository.ToggleUpvote(reportUpvote)

	if err != nil {
		return entities.ReportUpvote{}, "", err
	}

	return reportUpvote, status, nil
}
