package utils

import (
	"lawan-tambang-liar/constants"
	"net/http"
)

func ConvertResponseCode(err error) int {
	var badRequestErrors = []error{
		constants.ErrInvalidUsernameOrPassword,
		constants.ErrAllFieldsMustBeFilled,
		constants.ErrEmailAlreadyExist,
		constants.ErrUsernameAlreadyExist,
		constants.ErrMaxFileSize,
		constants.ErrMaxFileUpload,
		constants.ErrLimitAndPageMustBeFilled,
		constants.ErrIDMustBeFilled,
		constants.ErrReportNotFound,
		constants.ErrUserNotFound,
	}

	if contains(badRequestErrors, err) {
		return http.StatusBadRequest
	} else if err == constants.ErrUnauthorized {
		return http.StatusUnauthorized
	} else {
		return http.StatusInternalServerError
	}
}

func contains(slice []error, item error) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
