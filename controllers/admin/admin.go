package admin

import (
	"lawan-tambang-liar/controllers/admin/request"
	"lawan-tambang-liar/controllers/admin/response"
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUseCase entities.AdminUseCaseInterface
}

func NewAdminController(adminUseCase entities.AdminUseCaseInterface) *AdminController {
	return &AdminController{
		adminUseCase: adminUseCase,
	}
}

func (ac *AdminController) CreateAccount(c echo.Context) error {
	var adminRequest request.CreateAccount
	c.Bind(&adminRequest)

	admin, err := ac.adminUseCase.CreateAccount(adminRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.CreateAccountFromEntitiesToResponse(&admin)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Account", adminResponse))
}

func (ac *AdminController) Login(c echo.Context) error {
	var adminRequest request.Login
	c.Bind(&adminRequest)

	admin, err := ac.adminUseCase.Login(adminRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	JwtToken := new(http.Cookie)
	JwtToken.Name = "JwtToken"
	JwtToken.Value = admin.Token
	JwtToken.HttpOnly = true
	JwtToken.Secure = true
	JwtToken.Path = "/"
	JwtToken.Expires = time.Now().Add(time.Hour * 1)
	c.SetCookie(JwtToken)

	adminResponse := response.LoginFromEntitiesToResponse(&admin)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", adminResponse))
}
