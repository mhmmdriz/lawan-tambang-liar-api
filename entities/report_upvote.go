package entities

type ReportUpvote struct {
	ID       int `gorm:"primaryKey"`
	ReportID int
	UserID   int
	Report   Report `gorm:"foreignKey:ReportID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User     User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ReportUpvoteRepositoryInterface interface {
	ToggleUpvote(reportUpvote ReportUpvote) (ReportUpvote, string, error)
}

type ReportUpvoteUseCaseInterface interface {
	ToggleUpvote(reportID int, userID int) (ReportUpvote, string, error)
}
