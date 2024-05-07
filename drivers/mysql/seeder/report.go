package seeder

import (
	"errors"
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

func SeedReport(db *gorm.DB) {
	if err := db.First(&entities.Report{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create array of entities.Report
		reports := []entities.Report{
			{
				UserID:      1,
				Title:       "Report 1",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "1901",
				DistrictID:  "190102",
				Address:     "Pelabuhan Tanjung Gudang Belinyu",
				Files: []entities.ReportFile{
					{
						Path: "report_files/example1.jpg",
					},
				},
				Status: "pending",
			},
			{
				UserID:      2,
				Title:       "Report 2",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "1902",
				DistrictID:  "190204",
				Address:     "Jl. Raya Sijuk Desa Sijuk",
				Files: []entities.ReportFile{
					{
						Path: "report_files/example1.jpg",
					},
					{
						Path: "report_files/example2.jpg",
					},
				},
				Status: "verified",
			},
			{
				UserID:      3,
				Title:       "Report 3",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "1903",
				DistrictID:  "190304",
				Address:     "Pantai Batu Bedaun Desa Radjik",
				Files: []entities.ReportFile{
					{
						Path: "report_files/example1.jpg",
					},
					{
						Path: "report_files/example2.jpg",
					},
				},
				Status: "on progress",
			},
			{
				UserID:      4,
				Title:       "Report 4",
				Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				RegencyID:   "1904",
				DistrictID:  "190406",
				Address:     "Pantai Tanjung Berikat Desa Batu Beriga",
				Files: []entities.ReportFile{
					{
						Path: "report_files/example1.jpg",
					},
				},
				Status: "finished",
			},
		}

		if err := db.CreateInBatches(&reports, len(reports)).Error; err != nil {
			panic(err)
		}
	}
}
