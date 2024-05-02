package regency

import (
	"errors"
	"lawan-tambang-liar/entities"

	"gorm.io/gorm"
)

type RegencyRepo struct {
	DB *gorm.DB
}

func NewRegencyRepo(db *gorm.DB) *RegencyRepo {
	return &RegencyRepo{
		DB: db,
	}
}

func (r *RegencyRepo) AddRegenciesFromAPI(regencies []entities.Regency) error {
	if err := r.DB.First(&entities.Regency{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err := r.DB.CreateInBatches(regencies, len(regencies)).Error; err != nil {
			return err
		}
	} else {
		return errors.New("regency data already seeded in database")
	}

	return nil
}

func (r *RegencyRepo) GetRegencyIDs() ([]string, error) {
	var regencyIDs []string
	if err := r.DB.Model(&entities.Regency{}).Pluck("id", &regencyIDs).Error; err != nil {
		return nil, err
	}

	return regencyIDs, nil
}
