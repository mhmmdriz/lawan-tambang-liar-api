package seeder

import (
	"errors"
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

func SeedReportSolution(db *gorm.DB) {
	if err := db.First(&entities.ReportSolutionProcess{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create array of entities.ReportSolutionProcess
		reportSolutionProcesses := []entities.ReportSolutionProcess{
			{
				ReportID: 2,
				AdminID:  2,
				Message:  "Laporan ini sudah diverifikasi oleh admin kami. Terima kasih atas laporannya.",
				Status:   "verified",
			},
			{
				ReportID: 3,
				AdminID:  3,
				Message:  "Laporan ini sudah diverifikasi oleh admin kami. Terima kasih atas laporannya.",
				Status:   "verified",
			},
			{
				ReportID: 3,
				AdminID:  3,
				Message:  "Saat ini laporan ini sedang dalam proses penanganan.",
				Status:   "on progress",
				Files: []entities.ReportSolutionProcessFile{
					{
						Path: "report_solution_files/example1.jgp",
					},
				},
			},
			{
				ReportID: 4,
				AdminID:  4,
				Message:  "Laporan ini sudah diverifikasi oleh admin kami. Terima kasih atas laporannya.",
				Status:   "verified",
			},
			{
				ReportID: 4,
				AdminID:  4,
				Message:  "Saat ini laporan ini sedang dalam proses penanganan.",
				Status:   "on progress",
				Files: []entities.ReportSolutionProcessFile{
					{
						Path: "report_solution_files/example1.jgp",
					},
					{
						Path: "report_solution_files/example2.jgp",
					},
				},
			},
			{
				ReportID: 4,
				AdminID:  4,
				Message:  "Laporan ini sudah selesai ditangani oleh admin kami. Terima kasih atas laporannya.",
				Status:   "finished",
				Files: []entities.ReportSolutionProcessFile{
					{
						Path: "report_solution_files/example1.jgp",
					},
					{
						Path: "report_solution_files/example2.jgp",
					},
				},
			},
		}

		if err := db.CreateInBatches(&reportSolutionProcesses, len(reportSolutionProcesses)).Error; err != nil {
			panic(err)
		}
	}
}
