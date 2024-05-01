package main

import (
	"lawan-tambang-liar/config"
	regency_cl "lawan-tambang-liar/controllers/regency"
	regency_api "lawan-tambang-liar/drivers/indonesia_area_api/regency"
	"lawan-tambang-liar/drivers/mysql"
	regency_rp "lawan-tambang-liar/drivers/mysql/regency"
	"lawan-tambang-liar/routes"
	regency_uc "lawan-tambang-liar/usecases/regency"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.LoadEnv()
	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())

	e := echo.New()

	regencyAPI := regency_api.NewRegencyAPI()
	regencyRepo := regency_rp.NewRegencyRepo(DB)
	regencyUsecase := regency_uc.NewRegencyUsecase(regencyRepo, regencyAPI)
	RegencyController := regency_cl.NewRegencyController(regencyUsecase)

	routes := routes.RouteController{
		RegencyController: RegencyController,
	}

	routes.InitRoute(e)

	e.Start(":8080")
}
