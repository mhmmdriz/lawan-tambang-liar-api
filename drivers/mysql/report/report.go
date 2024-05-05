package report

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"time"

	"gorm.io/gorm"
)

type ReportRepo struct {
	DB *gorm.DB
}

func NewReportRepo(db *gorm.DB) *ReportRepo {
	return &ReportRepo{
		DB: db,
	}
}

func (r *ReportRepo) Create(report *entities.Report) error {
	if err := r.DB.Create(&report).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReportRepo) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sort_by string, sort_type string) ([]entities.Report, error) {
	var reports []entities.Report
	query := r.DB

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if sort_by != "" && sort_type != "" {
		query = query.Order(sort_by + " " + sort_type)
	} else if sort_by != "" && sort_type == "" {
		query = query.Order(sort_by + " DESC")
	} else if sort_by == "" && sort_type != "" {
		query = query.Order("created_at " + sort_type)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Limit(limit).Offset((page - 1) * limit).Preload("User").Preload("Regency").Preload("District").Preload("Files").Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *ReportRepo) Delete(report_id int, user_id int) (entities.Report, error) {
	var report entities.Report

	// Soft delete report by marking it as inactive
	if err := r.DB.First(&report, report_id).Error; err != nil {
		return entities.Report{}, constants.ErrReportNotFound
	}

	// Check if the user is the owner of the report
	if report.UserID != user_id {
		return entities.Report{}, constants.ErrUnauthorized
	}

	// Set the 'deleted_at' field to the current time
	report.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	// Update the report to mark it as deleted
	if err := r.DB.Save(&report).Error; err != nil {
		return entities.Report{}, constants.ErrInternalServerError
	}

	return report, nil
}
