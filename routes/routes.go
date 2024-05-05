package routes

import (
	"lawan-tambang-liar/controllers/admin"
	"lawan-tambang-liar/controllers/district"
	"lawan-tambang-liar/controllers/regency"
	"lawan-tambang-liar/controllers/report"
	"lawan-tambang-liar/controllers/user"
	"lawan-tambang-liar/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	RegencyController  *regency.RegencyController
	DistrictController *district.DistrictController
	UserController     *user.UserController
	AdminController    *admin.AdminController
	ReportController   *report.ReportController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	var jwtConfig = echojwt.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:   "cookie:JwtToken",
	}

	superAdmin := e.Group("/api/v1/super-admin")
	superAdmin.Use(echojwt.WithConfig(jwtConfig), middlewares.IsSuperAdmin)
	superAdmin.POST("/create-account", r.AdminController.CreateAccount)

	admin := e.Group("/api/v1/admin")
	admin.POST("/login", r.AdminController.Login)
	admin.Use(echojwt.WithConfig(jwtConfig), middlewares.IsAdmin)
	admin.POST("/seed-regency-db-from-api", r.RegencyController.SeedRegencyDBFromAPI)
	admin.POST("/seed-district-db-from-api", r.DistrictController.SeedDistrictDBFromAPI)
	admin.GET("/report", r.ReportController.GetPaginated)
	admin.GET("/report/:id", r.ReportController.GetByID)

	user := e.Group("/api/v1/user")
	user.POST("/register", r.UserController.Register)
	user.POST("/login", r.UserController.Login)
	user.Use(echojwt.WithConfig(jwtConfig), middlewares.IsUser)
	user.POST("/report", r.ReportController.Create)
	user.GET("/report", r.ReportController.GetPaginated)
	user.GET("/report/:id", r.ReportController.GetByID)
	user.DELETE("/report/:id", r.ReportController.Delete)

}
