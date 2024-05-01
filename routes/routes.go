package routes

import (
	"lawan-tambang-liar/controllers/district"
	"lawan-tambang-liar/controllers/regency"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	RegencyController  *regency.RegencyController
	DistrictController *district.DistrictController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/seed-regency-db-from-api", r.RegencyController.SeedRegencyDBFromAPI)
	e.POST("/seed-district-db-from-api", r.DistrictController.SeedDistrictDBFromAPI)
}
