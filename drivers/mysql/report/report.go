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

func (r *ReportRepo) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Report, error) {
	var reports []entities.Report
	query := r.DB

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if sortBy != "" && sortType != "" {
		query = query.Order(sortBy + " " + sortType)
	} else if sortBy != "" && sortType == "" {
		query = query.Order(sortBy + " DESC")
	} else if sortBy == "" && sortType != "" {
		query = query.Order("created_at " + sortType)
	} else {
		query = query.Order("created_at DESC")
	}

	if err := query.Limit(limit).Offset((page - 1) * limit).Preload("User").Preload("Regency").Preload("District").Preload("Files").Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *ReportRepo) GetByID(id int) (entities.Report, error) {
	var report entities.Report

	if err := r.DB.Preload("User").Preload("Regency").Preload("District").Preload("Files").First(&report, id).Error; err != nil {
		return entities.Report{}, constants.ErrReportNotFound
	}

	return report, nil
}

func (r *ReportRepo) Update(report entities.Report) (entities.Report, error) {
	// Check if the report exists
	var reportDB entities.Report

	if err := r.DB.First(&reportDB, report.ID).Error; err != nil {
		return entities.Report{}, constants.ErrReportNotFound
	}

	// Check if the user is the owner of the report
	if reportDB.UserID != report.UserID {
		return entities.Report{}, constants.ErrUnauthorized
	}

	reportDB.Title = report.Title
	reportDB.Description = report.Description
	reportDB.RegencyID = report.RegencyID
	reportDB.DistrictID = report.DistrictID
	reportDB.Address = report.Address

	// Update the report
	if err := r.DB.Save(&reportDB).Error; err != nil {
		return entities.Report{}, constants.ErrInternalServerError
	}

	report.Status = reportDB.Status

	return report, nil
}

func (r *ReportRepo) Delete(reportID int, userID int) (entities.Report, error) {
	var report entities.Report

	// Soft delete report by marking it as inactive
	if err := r.DB.First(&report, reportID).Error; err != nil {
		return entities.Report{}, constants.ErrReportNotFound
	}

	// Check if the user is the owner of the report
	if report.UserID != userID {
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

func (r *ReportRepo) AdminDelete(reportID int) (entities.Report, error) {
	var report entities.Report

	// Soft delete report by marking it as inactive
	if err := r.DB.First(&report, reportID).Error; err != nil {
		return entities.Report{}, constants.ErrReportNotFound
	}

	// Set the 'deleted_at' field to the current time
	report.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	// Update the report to mark it as deleted
	if err := r.DB.Save(&report).Error; err != nil {
		return entities.Report{}, constants.ErrInternalServerError
	}

	return report, nil
}

func (r *ReportRepo) UpdateStatus(reportID int, status string) error {
	var report entities.Report

	// Update the status of the report
	if err := r.DB.First(&report, reportID).Error; err != nil {
		return constants.ErrReportNotFound
	}

	report.Status = status

	if err := r.DB.Save(&report).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *ReportRepo) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var totalData int64
	var pagination entities.Pagination

	query := r.DB.Model(&entities.Report{})

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&totalData).Error; err != nil {
		return entities.Metadata{}, err
	}

	pagination.FirstPage = 1
	pagination.LastPage = (int(totalData) + limit - 1) / limit
	pagination.CurrentPage = page
	if pagination.CurrentPage == pagination.LastPage {
		pagination.TotalDataPerPage = int(totalData) - (pagination.LastPage-1)*limit
	} else {
		pagination.TotalDataPerPage = limit
	}

	if page > 1 {
		pagination.PrevPage = page - 1
	} else {
		pagination.PrevPage = 0
	}

	if page < pagination.LastPage {
		pagination.NextPage = page + 1
	} else {
		pagination.NextPage = 0
	}

	metadata := entities.Metadata{
		TotalData:  int(totalData),
		Pagination: pagination,
	}

	return metadata, nil
}

func (r *ReportRepo) IncreaseUpvote(reportID int) error {
	var report entities.Report

	// Increase the upvote count of the report
	if err := r.DB.First(&report, reportID).Error; err != nil {
		return constants.ErrReportNotFound
	}

	report.Upvotes++

	if err := r.DB.Save(&report).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *ReportRepo) DecreaseUpvote(reportID int) error {
	var report entities.Report

	// Decrease the upvote count of the report
	if err := r.DB.First(&report, reportID).Error; err != nil {
		return constants.ErrReportNotFound
	}

	report.Upvotes--

	if err := r.DB.Save(&report).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}
