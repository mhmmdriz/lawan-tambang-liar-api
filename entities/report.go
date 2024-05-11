package entities

import (
	"time"

	"gorm.io/gorm"
)

type Pagination struct {
	TotalDataPerPage int
	FirstPage        int
	LastPage         int
	CurrentPage      int
	NextPage         int
	PrevPage         int
}

type Metadata struct {
	TotalData  int
	Pagination Pagination
}

type Report struct {
	ID                      int
	UserID                  int
	Title                   string
	Description             string
	RegencyID               string `gorm:"type:varchar;size:191"`
	DistrictID              string `gorm:"type:varchar;size:191"`
	Address                 string
	Status                  string `gorm:"default:'pending';type:enum('pending', 'verified', 'on progress', 'finished', 'rejected')"`
	Upvotes                 int
	CreatedAt               time.Time               `gorm:"autoCreateTime"`
	UpdatedAt               time.Time               `gorm:"autoUpdateTime"`
	DeletedAt               gorm.DeletedAt          `gorm:"index"`
	User                    User                    `gorm:"foreignKey:UserID;references:ID"`
	Regency                 Regency                 `gorm:"foreignKey:RegencyID;references:ID"`
	District                District                `gorm:"foreignKey:DistrictID;references:ID"`
	Files                   []ReportFile            `gorm:"foreignKey:ReportID;references:ID"`
	ReportSolutionProcesses []ReportSolutionProcess `gorm:"foreignKey:ReportID;references:ID"`
}

type ReportRepositoryInterface interface {
	Create(report *Report) error
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]Report, error)
	GetByID(id int) (Report, error)
	Update(report Report) (Report, error)
	Delete(reportID int, userID int) (Report, error)
	AdminDelete(reportID int) (Report, error)
	UpdateStatus(reportID int, status string) error
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
	IncreaseUpvote(reportID int) error
	DecreaseUpvote(reportID int) error
}

type AIReportAPIInterface interface {
	GetChatCompletion(messages []map[string]string) (string, error)
}

type AdminReportRepositoryInterface interface {
	GetByID(id int) (Admin, error)
}

type ReportUseCaseInterface interface {
	Create(report *Report) (Report, error)
	GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]Report, error)
	GetByID(id int) (Report, error)
	Update(report Report) (Report, error)
	Delete(reportID int, userID int) (Report, error)
	AdminDelete(reportID int) (Report, error)
	UpdateStatus(reportID int, status string) error
	GetMetaData(limit int, page int, search string, filter map[string]interface{}) (Metadata, error)
	GetDistanceDuration(reportID int, userID int) (DistanceMatrix, error)
	GetDescriptionRecommendation(location string) (string, error)
	IncreaseUpvote(reportID int) error
	DecreaseUpvote(reportID int) error
}
