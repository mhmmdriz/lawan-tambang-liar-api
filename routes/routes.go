package routes

import (
	"lawan-tambang-liar/controllers/admin"
	"lawan-tambang-liar/controllers/district"
	"lawan-tambang-liar/controllers/regency"
	"lawan-tambang-liar/controllers/report"
	"lawan-tambang-liar/controllers/report_solution_process"
	"lawan-tambang-liar/controllers/user"
	"lawan-tambang-liar/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	RegencyController               *regency.RegencyController
	DistrictController              *district.DistrictController
	UserController                  *user.UserController
	AdminController                 *admin.AdminController
	ReportController                *report.ReportController
	ReportSolutionProcessController *report_solution_process.ReportSolutionProcessController
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
	superAdmin.DELETE("/delete-account/:id", r.AdminController.DeleteAccount)

	admin := e.Group("/api/v1/admin")
	admin.POST("/login", r.AdminController.Login)
	admin.Use(echojwt.WithConfig(jwtConfig), middlewares.IsAdmin)
	admin.GET("/:id", r.AdminController.GetByID)
	admin.POST("/seed-regency-db-from-api", r.RegencyController.SeedRegencyDBFromAPI)
	admin.POST("/seed-district-db-from-api", r.DistrictController.SeedDistrictDBFromAPI)
	admin.GET("/regencies", r.RegencyController.GetAll)
	admin.GET("/regencies/:id", r.RegencyController.GetByID)
	admin.GET("/districts", r.DistrictController.GetAll)
	admin.GET("/districts/:id", r.DistrictController.GetByID)
	admin.GET("/reports", r.ReportController.GetPaginated)
	admin.GET("/reports/:id", r.ReportController.GetByID)
	admin.DELETE("/reports/:id", r.ReportController.AdminDelete)
	admin.GET("/reports/:id/solutions", r.ReportSolutionProcessController.GetByReportID)
	admin.POST("/reports/:id/solutions/:action", r.ReportSolutionProcessController.Create)
	admin.DELETE("/reports/:id/solutions/:action/delete", r.ReportSolutionProcessController.Delete)
	admin.PUT("/reports/:id/solutions/:action/update", r.ReportSolutionProcessController.Update)

	user := e.Group("/api/v1/user")
	user.POST("/register", r.UserController.Register)
	user.POST("/login", r.UserController.Login)
	user.Use(echojwt.WithConfig(jwtConfig), middlewares.IsUser)
	user.GET("/regencies", r.RegencyController.GetAll)
	user.GET("/regencies/:id", r.RegencyController.GetByID)
	user.GET("/districts", r.DistrictController.GetAll)
	user.GET("/districts/:id", r.DistrictController.GetByID)
	user.POST("/reports", r.ReportController.Create)
	user.GET("/reports", r.ReportController.GetPaginated)
	user.GET("/reports/:id", r.ReportController.GetByID)
	user.DELETE("/reports/:id", r.ReportController.Delete)
	user.PUT("/reports/:id/update", r.ReportController.Update)

}
