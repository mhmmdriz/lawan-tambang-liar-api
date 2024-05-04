package constants

import "errors"

var (
	ErrInternalServerError       = errors.New("internal server error")
	ErrAllFieldsMustBeFilled     = errors.New("all fields must be filled")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrEmailAlreadyExist         = errors.New("email already exist")
	ErrUsernameAlreadyExist      = errors.New("username already exist")
	ErrUnauthorized              = errors.New("unauthorized")
	ErrMaxFileSize               = errors.New("max file size is 10MB")
	ErrMaxFileUpload             = errors.New("max file upload is 3")
)
