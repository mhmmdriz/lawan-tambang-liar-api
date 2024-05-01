package entities

type District struct {
	ID        string `json:"id" gorm:"primaryKey"`
	RegencyID int    `json:"regency_id"`
	Name      string `json:"name"`
}

type DistrictRepositoryInterface interface {
}

type DistrictUsecaseInterface interface {
	GetDistrictsAPI() ([]District, error)
}

type DistrictIndonesiaAreaAPIInterface interface {
	GetDistrictsAPI() ([]District, error)
}
