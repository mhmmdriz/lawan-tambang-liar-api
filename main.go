package main

import (
	"fmt"
	"lawan-tambang-liar/config"
	"lawan-tambang-liar/drivers/mysql"
	"lawan-tambang-liar/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.LoadEnv()
	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())
	e := echo.New()

	fmt.Println(DB)

	routes := routes.RouteController{}

	routes.InitRoute(e)
	e.Start(":8080")
}
