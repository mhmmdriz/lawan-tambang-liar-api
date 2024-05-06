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

func (r *DistrictRepo) GetAll(regencyID string) ([]entities.District, error) {
	var districts []entities.District
	if regencyID != "" {
		if err := r.DB.Where("regency_id = ?", regencyID).Find(&districts).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.DB.Find(&districts).Error; err != nil {
			return nil, err
		}
	}

	return districts, nil
}

func (r *DistrictRepo) GetByID(id string) (entities.District, error) {
	var district entities.District
	if err := r.DB.Where("id = ?", id).First(&district).Error; err != nil {
		return entities.District{}, err
	}

	return district, nil
}
