package main

import (
	"lawan-tambang-liar/config"
	admin_cl "lawan-tambang-liar/controllers/admin"
	district_cl "lawan-tambang-liar/controllers/district"
	regency_cl "lawan-tambang-liar/controllers/regency"
	report_cl "lawan-tambang-liar/controllers/report"
	report_solution_cl "lawan-tambang-liar/controllers/report_solution_process"
	report_upvote_cl "lawan-tambang-liar/controllers/report_upvote"
	user_cl "lawan-tambang-liar/controllers/user"
	ai_api "lawan-tambang-liar/drivers/ai_api"
	upload_file_gcs_api "lawan-tambang-liar/drivers/google_cloud_storage"
	"lawan-tambang-liar/drivers/google_maps_api"
	district_api "lawan-tambang-liar/drivers/indonesia_area_api/district"
	regency_api "lawan-tambang-liar/drivers/indonesia_area_api/regency"
	"lawan-tambang-liar/drivers/mysql"
	admin_rp "lawan-tambang-liar/drivers/mysql/admin"
	district_rp "lawan-tambang-liar/drivers/mysql/district"
	regency_rp "lawan-tambang-liar/drivers/mysql/regency"
	report_rp "lawan-tambang-liar/drivers/mysql/report"
	report_file_rp "lawan-tambang-liar/drivers/mysql/report_file"
	report_solution_rp "lawan-tambang-liar/drivers/mysql/report_solution_process"
	report_solution_file_rp "lawan-tambang-liar/drivers/mysql/report_solution_process_file"
	report_upvote_rp "lawan-tambang-liar/drivers/mysql/report_upvote"
	user_rp "lawan-tambang-liar/drivers/mysql/user"
	"lawan-tambang-liar/routes"
	admin_uc "lawan-tambang-liar/usecases/admin"
	district_uc "lawan-tambang-liar/usecases/district"
	regency_uc "lawan-tambang-liar/usecases/regency"
	report_uc "lawan-tambang-liar/usecases/report"
	report_file_uc "lawan-tambang-liar/usecases/report_file"
	report_solution_uc "lawan-tambang-liar/usecases/report_solution_process"
	report_solution_file_uc "lawan-tambang-liar/usecases/report_solution_process_file"
	report_upvote_uc "lawan-tambang-liar/usecases/report_upvote"
	user_uc "lawan-tambang-liar/usecases/user"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// For local development only
	// config.LoadEnv()

	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())

	gcs_credentials := os.Getenv("GCS_CREDENTIALS")
	gmaps_api_key := os.Getenv("GMAPS_API_KEY")

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

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

	aiAPI := ai_api.NewAIAPI()
	gmapsAPI := google_maps_api.NewGoogleMapsAPI(gmaps_api_key)
	reportFileRepo := report_file_rp.NewReportFileRepo(DB)
	uploadFileGCSAPI := upload_file_gcs_api.NewFileUploadAPI(gcs_credentials, "report_files/")
	reportFileUseCase := report_file_uc.NewReportFileUseCase(reportFileRepo, uploadFileGCSAPI)
	reportRepo := report_rp.NewReportRepo(DB)
	reportUsecase := report_uc.NewReportUseCase(reportRepo, adminRepo, gmapsAPI, aiAPI)
	ReportController := report_cl.NewReportController(reportUsecase, reportFileUseCase)

	reportSolutionFileRepo := report_solution_file_rp.NewReportSolutionProcessFileRepo(DB)
	uploadFileReportSolutionGCSAPI := upload_file_gcs_api.NewFileUploadAPI(gcs_credentials, "report_solution_files/")
	reportSolutionFileUseCase := report_solution_file_uc.NewReportSolutionProcessFileUsecase(reportSolutionFileRepo, uploadFileReportSolutionGCSAPI)
	reportSolutionRepo := report_solution_rp.NewReportSolutionProcessRepo(DB)
	reportSolutionUsecase := report_solution_uc.NewReportSolutionProcessUseCase(reportSolutionRepo, aiAPI)
	ReportSolutionProcessController := report_solution_cl.NewReportSolutionProcessController(reportUsecase, reportSolutionUsecase, reportSolutionFileUseCase)

	reportUpvoteRepo := report_upvote_rp.NewReportUpvoteRepo(DB)
	reportUpvoteUseCase := report_upvote_uc.NewReportUpvoteUseCase(reportUpvoteRepo)
	ReportUpvoteController := report_upvote_cl.NewReportUpvoteController(reportUsecase, reportUpvoteUseCase)

	routes := routes.RouteController{
		RegencyController:               RegencyController,
		DistrictController:              DistrictController,
		UserController:                  UserController,
		AdminController:                 AdminController,
		ReportController:                ReportController,
		ReportSolutionProcessController: ReportSolutionProcessController,
		ReportUpvoteController:          ReportUpvoteController,
	}

	routes.InitRoute(e)

	e.Start(":8080")
}
