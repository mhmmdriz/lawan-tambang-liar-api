package regency

import "lawan-tambang-liar/entities"

type RegencyUsecase struct {
	repository entities.RegencyRepositoryInterface
	api        entities.RegencyIndonesiaAreaAPIInterface
}

func NewRegencyUsecase(repository entities.RegencyRepositoryInterface, api entities.RegencyIndonesiaAreaAPIInterface) *RegencyUsecase {
	return &RegencyUsecase{
		repository: repository,
		api:        api,
	}
}

func (u *RegencyUsecase) SeedRegencyDBFromAPI() ([]entities.Regency, error) {
	regencies, err := u.api.GetRegenciesDataFromAPI()
	if err != nil {
		return nil, err
	}

	if err := u.repository.AddRegenciesFromAPI(regencies); err != nil {
		return nil, err
	}

	return regencies, nil
}
