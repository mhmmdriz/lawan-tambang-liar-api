package report

import (
	"lawan-tambang-liar/entities"

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
