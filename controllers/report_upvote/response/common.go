package response

import "lawan-tambang-liar/entities"

type CommonResponse struct {
	ID       int `json:"id"`
	ReportID int `json:"report_id"`
	UserID   int `json:"user_id"`
}

func FromEntitiesToResponse(reportUpvote *entities.ReportUpvote) *CommonResponse {
	return &CommonResponse{
		ID:       reportUpvote.ID,
		ReportID: reportUpvote.ReportID,
		UserID:   reportUpvote.UserID,
	}
}
