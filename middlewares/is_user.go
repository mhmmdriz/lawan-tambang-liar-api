package middlewares

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get JwtToken from authorization header
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": constants.ErrUnauthorized.Error(),
			})
		}

		// Get JWT Token from Authorization Header
		jwtToken := utils.GetToken(authorization)

		// Decode JWT Token Payload
		payload, err := utils.DecodePayload(jwtToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": constants.ErrInvalidJWT.Error(),
			})
		}

		role, ok := payload["role"].(string)
		if !ok || role != "user" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": constants.ErrUnauthorized.Error(),
			})
		}

		return next(c)
	}
}
