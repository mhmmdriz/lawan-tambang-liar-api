package district

import (
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/controllers/district/response"
	"lawan-tambang-liar/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DistrictController struct {
	districtUseCase entities.DistrictUseCaseInterface
	regencyUseCase  entities.RegencyUseCaseInterface
}

func NewDistrictController(districtUseCase entities.DistrictUseCaseInterface, regencyUseCase entities.RegencyUseCaseInterface) *DistrictController {
	return &DistrictController{
		districtUseCase: districtUseCase,
		regencyUseCase:  regencyUseCase,
	}
}

func (r *DistrictController) SeedDistrictDBFromAPI(c echo.Context) error {
	regencyIDs, err := r.regencyUseCase.GetRegencyIDs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	districts, err := r.districtUseCase.SeedDistrictDBFromAPI(regencyIDs)

	districts_response := []*response.DistrictSeedingResponse{}
	for _, district := range districts {
		districts_response = append(districts_response, response.FromUseCaseToResponse(&district))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Seeding Districts Table Data from API", districts_response))
}
