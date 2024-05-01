package regency

import (
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/controllers/regency/response"
	"lawan-tambang-liar/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegencyController struct {
	regencyUsecase entities.RegencyUsecaseInterface
}

func NewRegencyController(regencyUsecase entities.RegencyUsecaseInterface) *RegencyController {
	return &RegencyController{
		regencyUsecase: regencyUsecase,
	}
}

func (r *RegencyController) SeedRegencyDBFromAPI(c echo.Context) error {
	regencies, err := r.regencyUsecase.SeedRegencyDBFromAPI()

	regencies_response := []*response.RegencySeedingResponse{}
	for _, regency := range regencies {
		regencies_response = append(regencies_response, response.FromUseCaseToResponse(&regency))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Seeding Regencies Table Data from API", regencies_response))
}
