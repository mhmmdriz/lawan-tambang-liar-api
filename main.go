package main

import (
	"lawan-tambang-liar/config"
	admin_cl "lawan-tambang-liar/controllers/admin"
	district_cl "lawan-tambang-liar/controllers/district"
	regency_cl "lawan-tambang-liar/controllers/regency"
	report_cl "lawan-tambang-liar/controllers/report"
	user_cl "lawan-tambang-liar/controllers/user"
	upload_file_gcs_api "lawan-tambang-liar/drivers/google_cloud_storage"
	district_api "lawan-tambang-liar/drivers/indonesia_area_api/district"
	regency_api "lawan-tambang-liar/drivers/indonesia_area_api/regency"
	"lawan-tambang-liar/drivers/mysql"
	admin_rp "lawan-tambang-liar/drivers/mysql/admin"
	district_rp "lawan-tambang-liar/drivers/mysql/district"
	regency_rp "lawan-tambang-liar/drivers/mysql/regency"
	report_rp "lawan-tambang-liar/drivers/mysql/report"
	report_file_rp "lawan-tambang-liar/drivers/mysql/report_file"
	user_rp "lawan-tambang-liar/drivers/mysql/user"
	"lawan-tambang-liar/routes"
	admin_uc "lawan-tambang-liar/usecases/admin"
	district_uc "lawan-tambang-liar/usecases/district"
	regency_uc "lawan-tambang-liar/usecases/regency"
	report_uc "lawan-tambang-liar/usecases/report"
	report_file_uc "lawan-tambang-liar/usecases/report_file"
	user_uc "lawan-tambang-liar/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())

	e := echo.New()

	regencyAPI := regency_api.NewRegencyAPI()
	regencyRepo := regency_rp.NewRegencyRepo(DB)
	regencyUsecase := regency_uc.NewRegencyUsecase(regencyRepo, regencyAPI)
	RegencyController := regency_cl.NewRegencyController(regencyUsecase)

	districtAPI := district_api.NewDistrictAPI()
	districtRepo := district_rp.NewDistrictRepo(DB)
	districtUsecase := district_uc.NewDistrictUseCase(districtRepo, districtAPI)
	DistrictController := district_cl.NewDistrictController(districtUsecase, regencyUsecase)

	userRepo := user_rp.NewUserRepo(DB)
	userUsecase := user_uc.NewUserUseCase(userRepo)
	UserController := user_cl.NewUserController(userUsecase)

	adminRepo := admin_rp.NewAdminRepo(DB)
	adminUsecase := admin_uc.NewAdminUseCase(adminRepo)
	AdminController := admin_cl.NewAdminController(adminUsecase)

	reportFileRepo := report_file_rp.NewReportFileRepo(DB)
	uploadFileGCSAPI := upload_file_gcs_api.NewFileUploadAPI("report_files/")
	reportFileUseCase := report_file_uc.NewReportFileUseCase(reportFileRepo, uploadFileGCSAPI)
	reportRepo := report_rp.NewReportRepo(DB)
	reportUsecase := report_uc.NewReportUseCase(reportRepo)
	ReportController := report_cl.NewReportController(reportUsecase, reportFileUseCase)

	routes := routes.RouteController{
		RegencyController:  RegencyController,
		DistrictController: DistrictController,
		UserController:     UserController,
		AdminController:    AdminController,
		ReportController:   ReportController,
	}

	routes.InitRoute(e)

	e.Start(":8080")
}
