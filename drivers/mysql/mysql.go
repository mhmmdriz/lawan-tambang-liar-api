package mysql

import (
	"errors"
	"fmt"
	"lawan-tambang-liar/drivers/indonesia_area_api/regency"
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

	Migration(db)
	SeedRegencyFromAPI(db, regencyAPI)

	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&entities.Regency{})
	db.AutoMigrate(&entities.District{})
}

func SeedRegencyFromAPI(db *gorm.DB, api entities.RegencyIndonesiaAreaAPIInterface) {
	if err := db.First(&entities.Regency{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		regencies, err := api.GetRegenciesDataFromAPI()
		if err != nil {
			panic(err)
		}

		if err := db.CreateInBatches(regencies, len(regencies)).Error; err != nil {
			panic(err)
		}
	}
}
