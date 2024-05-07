package utils

import (
	"lawan-tambang-liar/constants"

	"github.com/labstack/echo/v4"
)

func GetUserIDFromJWT(c echo.Context) (int, error) {
	// Get jwt from cookie
	jwt, err := c.Cookie("JwtToken")
	if err != nil {
		return 0, constants.ErrUnauthorized
	}
	jwt_payload, err := DecodePayload(jwt.Value)
	if err != nil {
		return 0, constants.ErrInternalServerError
	}

	// Get user id from jwt payload
	user_id, ok := jwt_payload["id"].(float64)
	if !ok {
		return 0, constants.ErrInternalServerError
	}

	return int(user_id), nil
}
