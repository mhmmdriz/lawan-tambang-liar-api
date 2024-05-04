package report

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/controllers/report/request"
	response_report "lawan-tambang-liar/controllers/report/response"
	response_report_file "lawan-tambang-liar/controllers/report_file/response"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportController struct {
	reportUseCase     entities.ReportUseCaseInterface
	reportFileUseCase entities.ReportFileUseCaseInterface
}

func NewReportController(reportUseCase entities.ReportUseCaseInterface, reportFileUseCase entities.ReportFileUseCaseInterface) *ReportController {
	return &ReportController{
		reportUseCase:     reportUseCase,
		reportFileUseCase: reportFileUseCase,
	}
}

func (rc *ReportController) Create(c echo.Context) error {
	user_id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var reportRequest request.Create
	c.Bind(&reportRequest)
	reportRequest.UserID = user_id

	// Parse form-data multipart
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	// Mengambil semua file yang diunggah
	files := form.File["files"]

	if len(files) > 3 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrMaxFileUpload.Error()))
	}

	// Count total file size
	totalFileSize := 0
	for _, file := range files {
		totalFileSize += int(file.Size)
	}

	if totalFileSize > 10*1024*1024 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constants.ErrMaxFileSize.Error()))
	}

	report, err1 := rc.reportUseCase.Create(reportRequest.ToEntities())
	if err1 != nil {
		return c.JSON(utils.ConvertResponseCode(err1), base.NewErrorResponse(err1.Error()))
	}

	reportResponse := response_report.CreateFromEntitiesToResponse(&report)

	reportFile, err3 := rc.reportFileUseCase.Create(files, report.ID)
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err3.Error()))
	}

	reportFileResponses := []*response_report_file.ReportFile{}
	for _, rf := range reportFile {
		reportFileResponses = append(reportFileResponses, response_report_file.CreateFromEntitiesToResponse(&rf))
	}

	reportResponse.Files = reportFileResponses

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Report", reportResponse))
}
