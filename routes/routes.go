package routes

import (
	"lawan-tambang-liar/controllers/regency"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	RegencyController *regency.RegencyController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/seed-regency-db-from-api", r.RegencyController.SeedRegencyDBFromAPI)
}
