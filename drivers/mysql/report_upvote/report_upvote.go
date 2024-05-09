package report_upvote

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

type ReportUpvoteRepo struct {
	DB *gorm.DB
}

func NewReportUpvoteRepo(db *gorm.DB) *ReportUpvoteRepo {
	return &ReportUpvoteRepo{
		DB: db,
	}
}

func (r *ReportUpvoteRepo) ToggleUpvote(reportUpvote entities.ReportUpvote) (entities.ReportUpvote, string, error) {
	var existingReportUpvote entities.ReportUpvote

	if err := r.DB.Where("report_id = ? AND user_id = ?", reportUpvote.ReportID, reportUpvote.UserID).First(&existingReportUpvote).Error; err != nil {
		if err := r.DB.Create(&reportUpvote).Error; err != nil {
			return entities.ReportUpvote{}, "", constants.ErrReportNotFound
		}
		return reportUpvote, "upvoted", nil
	}

	if err := r.DB.Delete(&existingReportUpvote).Error; err != nil {
		return entities.ReportUpvote{}, "", constants.ErrInternalServerError
	}

	return existingReportUpvote, "downvoted", nil
}
