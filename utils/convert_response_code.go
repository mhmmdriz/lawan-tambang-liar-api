package utils

import (
	"lawan-tambang-liar/constants"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch err {
	case constants.ErrInvalidUsernameOrPassword:
		return http.StatusBadRequest
	case constants.ErrAllFieldsMustBeFilled:
		return http.StatusBadRequest
	case constants.ErrEmailAlreadyExist:
		return http.StatusBadRequest
	case constants.ErrUsernameAlreadyExist:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
