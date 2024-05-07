package user

import (
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/controllers/user/request"
	"lawan-tambang-liar/controllers/user/response"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"net/http"
	"strconv"
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
	var userRequest request.Register
	c.Bind(&userRequest)

	user, err := uc.userUseCase.Register(userRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	userResponse := response.RegisterFromEntitiesToResponse(&user)

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Register", userResponse))
}

func (uc *UserController) Login(c echo.Context) error {
	var userRequest request.Login
	c.Bind(&userRequest)

	user, err := uc.userUseCase.Login(userRequest.ToEntities())
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

	userResponse := response.LoginFromEntitiesToResponse(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", userResponse))
}

func (uc *UserController) GetAll(c echo.Context) error {
	users, err := uc.userUseCase.GetAll()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var usersResponse []response.Get
	for _, user := range users {
		usersResponse = append(usersResponse, *response.GetFromEntitiesToResponse(&user))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Users", usersResponse))
}

func (uc *UserController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := uc.userUseCase.GetByID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.GetFromEntitiesToResponse(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get User By ID", userResponse))
}

func (uc *UserController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := uc.userUseCase.Delete(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.DeleteFromEntitiesToResponse(&user)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete User", userResponse))
}
