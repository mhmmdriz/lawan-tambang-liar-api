package user

import (
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/controllers/user/response"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase entities.UserUseCaseInterface
}

func NewUserController(userUseCase entities.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var userRegister entities.User
	c.Bind(&userRegister)

	user, err := uc.userUseCase.Register(&userRegister)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	userResponse := response.FromUseCaseToRegisterResponse(&user)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Register", userResponse))
}

func (uc *UserController) Login(c echo.Context) error {
	var userLogin entities.User
	c.Bind(&userLogin)

	user, err := uc.userUseCase.Login(&userLogin)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	JwtToken := new(http.Cookie)
	JwtToken.Name = "JwtToken"
	JwtToken.Value = user.Token
	JwtToken.HttpOnly = true
	JwtToken.Secure = true
	JwtToken.Path = "/"
	JwtToken.Expires = time.Now().Add(time.Hour * 1)
	c.SetCookie(JwtToken)

	userResponse := response.FromUseCaseToLoginResponse(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", userResponse))
}
