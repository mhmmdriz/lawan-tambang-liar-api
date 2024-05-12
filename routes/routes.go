package routes

import (
	"lawan-tambang-liar/controllers/admin"
	"lawan-tambang-liar/controllers/district"
	"lawan-tambang-liar/controllers/regency"
	"lawan-tambang-liar/controllers/report"
	"lawan-tambang-liar/controllers/report_solution_process"
	"lawan-tambang-liar/controllers/report_upvote"
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
	ReportUpvoteController          *report_upvote.ReportUpvoteController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	var jwt = echojwt.JWT([]byte(os.Getenv("JWT_SECRET")))

	superAdmin := e.Group("/api/v2/super-admin")
	superAdmin.POST("/login", r.AdminController.Login)
	superAdmin.Use(jwt, middlewares.IsSuperAdmin)
	superAdmin.GET("/admin-accounts", r.AdminController.GetAll)
	superAdmin.GET("/admin-accounts/:id", r.AdminController.GetByID)
	superAdmin.POST("/admin-accounts", r.AdminController.CreateAccount)
	superAdmin.DELETE("/admin-accounts/:id", r.AdminController.DeleteAccount)
	superAdmin.PUT("/admin-accounts/:id/reset-password", r.AdminController.ResetPassword)
	superAdmin.PUT("/change-password", r.AdminController.ChangePassword)

	admin := e.Group("/api/v2/admin")
	admin.POST("/login", r.AdminController.Login)
	admin.Use(jwt, middlewares.IsAdmin)
	admin.PUT("/reset-password", r.AdminController.ResetPassword)
	admin.PUT("/change-password", r.AdminController.ChangePassword)
	admin.POST("/seed-regency-db-from-api", r.RegencyController.SeedRegencyDBFromAPI)
	admin.POST("/seed-district-db-from-api", r.DistrictController.SeedDistrictDBFromAPI)
	admin.GET("/regencies", r.RegencyController.GetAll)
	admin.GET("/regencies/:id", r.RegencyController.GetByID)
	admin.GET("/districts", r.DistrictController.GetAll)
	admin.GET("/districts/:id", r.DistrictController.GetByID)
	admin.GET("/reports", r.ReportController.GetPaginated)
	admin.GET("/reports/:id", r.ReportController.GetByID)
	admin.GET("/reports/:id/distance-duration", r.ReportController.GetDistanceDuration)
	admin.DELETE("/reports/:id", r.ReportController.AdminDelete)
	admin.GET("/reports/:id/solutions", r.ReportSolutionProcessController.GetByReportID)
	admin.POST("/reports/:id/solutions/:action", r.ReportSolutionProcessController.Create)
	admin.DELETE("/reports/:id/solutions/:action", r.ReportSolutionProcessController.Delete)
	admin.PUT("/reports/:id/solutions/:action", r.ReportSolutionProcessController.Update)
	admin.GET("/reports/message-recommendation/:action", r.ReportSolutionProcessController.GetMessageRecommendation)
	admin.GET("/user-accounts", r.UserController.GetAll)
	admin.GET("/user-accounts/:id", r.UserController.GetByID)
	admin.DELETE("/user-accounts/:id", r.UserController.Delete)
	admin.PUT("/user-accounts/:id/reset-password", r.UserController.ResetPassword)

	user := e.Group("/api/v2/user")
	user.POST("/register", r.UserController.Register)
	user.POST("/login", r.UserController.Login)
	user.Use(jwt, middlewares.IsUser)
	user.PUT("/change-password", r.UserController.ChangePassword)
	user.GET("/regencies", r.RegencyController.GetAll)
	user.GET("/regencies/:id", r.RegencyController.GetByID)
	user.GET("/districts", r.DistrictController.GetAll)
	user.GET("/districts/:id", r.DistrictController.GetByID)
	user.POST("/reports", r.ReportController.Create)
	user.GET("/reports", r.ReportController.GetPaginated)
	user.GET("/reports/:id", r.ReportController.GetByID)
	user.DELETE("/reports/:id", r.ReportController.Delete)
	user.PUT("/reports/:id", r.ReportController.Update)
	user.GET("/reports/description-recommendation", r.ReportController.GetDescriptionRecommendation)
	user.POST("/reports/:id/upvote", r.ReportUpvoteController.ToggleUpvote)
	user.GET("/reports/:id/solutions", r.ReportSolutionProcessController.GetByReportID)

}
