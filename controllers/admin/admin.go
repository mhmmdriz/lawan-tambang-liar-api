package admin

import (
	"lawan-tambang-liar/controllers/admin/request"
	"lawan-tambang-liar/controllers/admin/response"
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"net/http"
	"strconv"
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

func (ac *AdminController) GetAll(c echo.Context) error {
	admins, err := ac.adminUseCase.GetAll()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var adminsResponse []*response.Get
	for _, admin := range admins {
		adminsResponse = append(adminsResponse, response.GetFromEntitiesToResponse(&admin))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Admin", adminsResponse))
}

func (ac *AdminController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	admin, err := ac.adminUseCase.GetByID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.GetFromEntitiesToResponse(&admin)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Admin", adminResponse))
}

func (ac *AdminController) DeleteAccount(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	admin, err := ac.adminUseCase.DeleteAccount(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.DeleteFromEntitiesToResponse(&admin)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Admin", adminResponse))
}

func (ac *AdminController) ResetPassword(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	admin, err := ac.adminUseCase.ResetPassword(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	adminResponse := response.PasswordFromEntitiesToResponse(&admin)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Reset Password", adminResponse))
}

func (ac *AdminController) ChangePassword(c echo.Context) error {
	admin_id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	admin, err := ac.adminUseCase.ChangePassword(admin_id, c.FormValue("new_password"))
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	adminResponse := response.PasswordFromEntitiesToResponse(&admin)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Change Password", adminResponse))
}
