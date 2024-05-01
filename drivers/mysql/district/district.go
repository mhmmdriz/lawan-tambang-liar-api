package district

import (
	"errors"
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

type DistrictRepo struct {
	DB *gorm.DB
}

func NewDistrictRepo(db *gorm.DB) *DistrictRepo {
	return &DistrictRepo{
		DB: db,
	}
}

func (r *DistrictRepo) AddDistrictsFromAPI(districts []entities.District) error {
	if err := r.DB.First(&entities.District{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err := r.DB.CreateInBatches(districts, len(districts)).Error; err != nil {
			return err
		}
	} else {
		return errors.New("district data already seeded in database")
	}

	return nil
}
