package mysql

import (
	"fmt"
	"lawan-tambang-liar/drivers/indonesia_area_api/district"
	"lawan-tambang-liar/drivers/indonesia_area_api/regency"
	"lawan-tambang-liar/drivers/mysql/seeder"
	"lawan-tambang-liar/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func ConnectDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	regencyAPI := regency.NewRegencyAPI()
	districtAPI := district.NewDistrictAPI()

	Migration(db)
	seeder.SeedRegencyFromAPI(db, regencyAPI)
	seeder.SeedDistrictFromAPI(db, districtAPI)
	seeder.SeedAdmin(db)
	seeder.SeedUser(db)

	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&entities.Regency{})
	db.AutoMigrate(&entities.District{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Admin{})
	db.AutoMigrate(&entities.Report{})
	db.AutoMigrate(&entities.ReportFile{})
	db.AutoMigrate(&entities.ReportSolutionProcess{})
	db.AutoMigrate(&entities.ReportSolutionProcessFile{})
	db.AutoMigrate(&entities.ReportUpvote{})
}
