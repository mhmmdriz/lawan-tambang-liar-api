package entities

type District struct {
	ID        string `json:"id" gorm:"primaryKey"`
	RegencyID string `json:"regency_id" gorm:"foreignKey:RegencyID;references:ID;type:varchar;size:191"`
	Name      string `json:"name"`
}

type DistrictRepositoryInterface interface {
	AddDistrictsFromAPI(districts []District) error
}

type DistrictIndonesiaAreaAPIInterface interface {
	GetDistrictsDataFromAPI([]string) ([]District, error)
}

type DistrictUseCaseInterface interface {
	SeedDistrictDBFromAPI([]string) ([]District, error)
}
