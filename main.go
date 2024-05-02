package main

import (
	"lawan-tambang-liar/config"
	district_cl "lawan-tambang-liar/controllers/district"
	regency_cl "lawan-tambang-liar/controllers/regency"
	user_cl "lawan-tambang-liar/controllers/user"
	district_api "lawan-tambang-liar/drivers/indonesia_area_api/district"
	regency_api "lawan-tambang-liar/drivers/indonesia_area_api/regency"
	"lawan-tambang-liar/drivers/mysql"
	district_rp "lawan-tambang-liar/drivers/mysql/district"
	regency_rp "lawan-tambang-liar/drivers/mysql/regency"
	user_rp "lawan-tambang-liar/drivers/mysql/user"
	"lawan-tambang-liar/routes"
	district_uc "lawan-tambang-liar/usecases/district"
	regency_uc "lawan-tambang-liar/usecases/regency"
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

	routes := routes.RouteController{
		RegencyController:  RegencyController,
		DistrictController: DistrictController,
		UserController:     UserController,
	}

	routes.InitRoute(e)

	e.Start(":8080")
}
