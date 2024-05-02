package admin

import (
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
	var admin entities.Admin
	c.Bind(&admin)

	admin, err := ac.adminUseCase.CreateAccount(&admin)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.FromUseCaseToCreateAccountResponse(&admin)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Create Account", adminResponse))
}

func (ac *AdminController) Login(c echo.Context) error {
	var admin entities.Admin
	c.Bind(&admin)

	admin, err := ac.adminUseCase.Login(&admin)
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

	adminResponse := response.FromUseCaseToLoginResponse(&admin)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", adminResponse))
}
