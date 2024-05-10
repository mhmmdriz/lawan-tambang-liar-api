package constants

import "errors"

var (
	ErrInternalServerError           = errors.New("internal server error")
	ErrAllFieldsMustBeFilled         = errors.New("all fields must be filled")
	ErrInvalidUsernameOrPassword     = errors.New("invalid username or password")
	ErrEmailAlreadyExist             = errors.New("email already exist")
	ErrUsernameAlreadyExist          = errors.New("username already exist")
	ErrUnauthorized                  = errors.New("unauthorized")
	ErrMaxFileSize                   = errors.New("max file size is 10MB")
	ErrMaxFileUpload                 = errors.New("max file upload is 3")
	ErrLimitAndPageMustBeFilled      = errors.New("limit and page must be filled")
	ErrIDMustBeFilled                = errors.New("id must be filled")
	ErrReportNotFound                = errors.New("report not found")
	ErrActionNotFound                = errors.New("action not found")
	ErrReportSolutionProcessNotFound = errors.New("report solution process not found")
	ErrUserNotFound                  = errors.New("user not found")
	ErrAdminNotFound                 = errors.New("admin not found")
	ErrSuperAdminCannotBeDeleted     = errors.New("super admin cannot be deleted")
	ErrInvalidJWT                    = errors.New("invalid jwt")
)
