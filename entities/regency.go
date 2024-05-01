package entities

type Regency struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Districts []District `json:"districts" gorm:"foreignKey:RegencyID;references:ID"`
}

type RegencyRepositoryInterface interface {
	AddRegenciesFromAPI(regencies []Regency) error
	GetRegencyIDs() ([]string, error)
}

type RegencyIndonesiaAreaAPIInterface interface {
	GetRegenciesDataFromAPI() ([]Regency, error)
}

type RegencyUseCaseInterface interface {
	SeedRegencyDBFromAPI() ([]Regency, error)
	GetRegencyIDs() ([]string, error)
}
