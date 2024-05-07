package report_solution_process

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/entities"
	"time"

	"gorm.io/gorm"
)

type ReportSolutionProcessRepo struct {
	DB *gorm.DB
}

func NewReportSolutionProcessRepo(db *gorm.DB) *ReportSolutionProcessRepo {
	return &ReportSolutionProcessRepo{
		DB: db,
	}
}

func (r *ReportSolutionProcessRepo) Create(reportSolutionProcess *entities.ReportSolutionProcess) error {
	if err := r.DB.Create(&reportSolutionProcess).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReportSolutionProcessRepo) GetByReportID(reportID int) ([]entities.ReportSolutionProcess, error) {
	var reportSolutionProcesses []entities.ReportSolutionProcess

	if err := r.DB.Where("report_id = ?", reportID).Find(&reportSolutionProcesses).Error; err != nil {
		return nil, err
	}

	return reportSolutionProcesses, nil
}

func (r *ReportSolutionProcessRepo) Delete(reportID int, reportSolutionProcessStatus string) (entities.ReportSolutionProcess, error) {
	var reportSolutionProcess entities.ReportSolutionProcess

	// Soft delete
	if err := r.DB.Where("report_id = ? && status = ?", reportID, reportSolutionProcessStatus).First(&reportSolutionProcess).Error; err != nil {
		return entities.ReportSolutionProcess{}, err
	}

	reportSolutionProcess.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := r.DB.Save(&reportSolutionProcess).Error; err != nil {
		return entities.ReportSolutionProcess{}, err
	}

	return reportSolutionProcess, nil
}

func (r *ReportSolutionProcessRepo) Update(reportSolutionProcess entities.ReportSolutionProcess) (entities.ReportSolutionProcess, error) {
	// Check if the report solution process exists
	var reportSolutionDB entities.ReportSolutionProcess

	if err := r.DB.Where("report_id = ? && status = ?", reportSolutionProcess.ReportID, reportSolutionProcess.Status).First(&reportSolutionDB).Error; err != nil {
		return entities.ReportSolutionProcess{}, constants.ErrReportSolutionProcessNotFound
	}

	reportSolutionDB.Message = reportSolutionProcess.Message
	reportSolutionDB.AdminID = reportSolutionProcess.AdminID

	if err := r.DB.Save(&reportSolutionDB).Error; err != nil {
		return entities.ReportSolutionProcess{}, constants.ErrInternalServerError
	}

	reportSolutionProcess.ID = reportSolutionDB.ID
	reportSolutionProcess.Status = reportSolutionDB.Status

	return reportSolutionProcess, nil
}
