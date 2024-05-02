package routes

import (
	"lawan-tambang-liar/controllers/admin"
	"lawan-tambang-liar/controllers/district"
	"lawan-tambang-liar/controllers/regency"
	"lawan-tambang-liar/controllers/user"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	RegencyController  *regency.RegencyController
	DistrictController *district.DistrictController
	UserController     *user.UserController
	AdminController    *admin.AdminController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/api/v1/seed-regency-db-from-api", r.RegencyController.SeedRegencyDBFromAPI)
	e.POST("/api/v1/seed-district-db-from-api", r.DistrictController.SeedDistrictDBFromAPI)

	e.POST("/api/v1/user/register", r.UserController.Register)
	e.POST("/api/v1/user/login", r.UserController.Login)

	e.POST("/api/v1/admin/create-account", r.AdminController.CreateAccount)
	e.POST("/api/v1/admin/login", r.AdminController.Login)

	jwtAuth := e.Group("/api/v1")
	jwtAuth.Use(echojwt.WithConfig(echojwt.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:   "cookie:JwtToken",
	}))

	jwtAuth.GET("/test", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})
}
