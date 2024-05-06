package district

import "lawan-tambang-liar/entities"

type DistrictUseCase struct {
	repository entities.DistrictRepositoryInterface
	api        entities.DistrictIndonesiaAreaAPIInterface
}

func NewDistrictUseCase(repository entities.DistrictRepositoryInterface, api entities.DistrictIndonesiaAreaAPIInterface) *DistrictUseCase {
	return &DistrictUseCase{
		repository: repository,
		api:        api,
	}
}

func (u *DistrictUseCase) SeedDistrictDBFromAPI(regencyIDs []string) ([]entities.District, error) {
	districts, err := u.api.GetDistrictsDataFromAPI(regencyIDs)
	if err != nil {
		return nil, err
	}

	if err := u.repository.AddDistrictsFromAPI(districts); err != nil {
		return nil, err
	}

	return districts, nil
}

func (u *DistrictUseCase) GetAll(regencyID string) ([]entities.District, error) {
	districts, err := u.repository.GetAll(regencyID)
	if err != nil {
		return nil, err
	}

	return districts, nil
}

func (u *DistrictUseCase) GetByID(id string) (entities.District, error) {
	district, err := u.repository.GetByID(id)
	if err != nil {
		return entities.District{}, err
	}

	return district, nil
}
