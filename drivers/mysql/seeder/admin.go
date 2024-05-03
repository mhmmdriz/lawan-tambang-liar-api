package seeder

import (
	"errors"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	if err := db.First(&entities.Admin{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create array of entities.Admin
		hash, _ := utils.HashPassword("admin")
		admins := []entities.Admin{
			{
				Username:        "super_admin",
				Password:        hash,
				Email:           "super_admin@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    true,
				RegencyID:       "1971",
				DistrictID:      "197102",
				Address:         "Jl Balai, Gedung Nasional",
			},
			{
				Username:        "admin_bangka",
				Password:        hash,
				Email:           "admin_bangka@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
				RegencyID:       "1901",
				DistrictID:      "190101",
				Address:         "Jl. Jend. Sudirman No. 1, Sri Menanti",
			},
			{
				Username:        "admin_belitung",
				Password:        hash,
				Email:           "admin_belitung@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
				RegencyID:       "1902",
				DistrictID:      "190201",
				Address:         "Jl. Merdeka 16, Tj. Pandan",
			},
			{
				Username:        "admin_bangka_selatan",
				Password:        hash,
				Email:           "admin_bangka_selatan@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
				RegencyID:       "1903",
				DistrictID:      "190301",
				Address:         "Jl. Jend. Sudirman 16, Toboali",
			},
			{
				Username:        "admin_bangka_tengah",
				Password:        hash,
				Email:           "admin_bangka_tengah@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
				RegencyID:       "1904",
				DistrictID:      "190401",
				Address:         "Jl. Namang - Koba 12, Koba",
			},
			{
				Username:        "admin_bangka_barat",
				Password:        hash,
				Email:           "admin_bangka_barat@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
				RegencyID:       "1905",
				DistrictID:      "190501",
				Address:         "Jl. Jend. Sudirman 16, Muntok",
			},
			{
				Username:        "admin_belitung_timur",
				Password:        hash,
				Email:           "admin@gmail.com",
				TelephoneNumber: "081234567890",
				IsSuperAdmin:    false,
				RegencyID:       "1906",
				DistrictID:      "190601",
				Address:         "Jl. Raya Manggar-Gantung No.10, Manggar",
			},
		}

		if err := db.CreateInBatches(&admins, len(admins)).Error; err != nil {
			panic(err)
		}
	}
}
