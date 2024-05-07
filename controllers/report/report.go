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
	"strconv"

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
		reportFileResponses = append(reportFileResponses, response_report_file.FromEntitiesToResponse(&rf))
	}

	reportResponse.Files = reportFileResponses

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Report", reportResponse))
}

func (rc *ReportController) GetPaginated(c echo.Context) error {
	// Get limit, page, search query, and filter from query params
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	filter_district, _ := strconv.Atoi(c.QueryParam("district"))
	filter_regency, _ := strconv.Atoi(c.QueryParam("regency"))
	filter_status := c.QueryParam("status")
	filter := map[string]interface{}{}
	if filter_district == 0 && filter_regency == 0 && filter_status == "" {
		filter = nil
	} else {
		if filter_district != 0 {
			filter["district_id"] = filter_district
		}
		if filter_regency != 0 {
			filter["regency_id"] = filter_regency
		}
		if filter_status != "" {
			filter["status"] = filter_status
		}
	}

	sort_by := c.QueryParam("sort_by")
	sort_type := c.QueryParam("sort_type")

	reports, err := rc.reportUseCase.GetPaginated(limit, page, search, filter, sort_by, sort_type)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportResponses := []*response_report.GetPaginate{}
	var reportResponse *response_report.GetPaginate
	for _, report := range reports {
		reportResponse = response_report.GetPaginateFromEntitiesToResponse(&report)
		reportResponses = append(reportResponses, reportResponse)
	}

	metaData, err := rc.reportUseCase.GetMetaData(limit, page, search, filter)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	metaDataResponse := base.NewMetadata(metaData.TotalData, metaData.Pagination.TotalDataPerPage, metaData.Pagination.FirstPage, metaData.Pagination.LastPage, metaData.Pagination.CurrentPage, metaData.Pagination.NextPage, metaData.Pagination.PrevPage)

	return c.JSON(http.StatusOK, base.NewSuccessResponseWithMetadata("Success Get Reports", reportResponses, *metaDataResponse))
}

func (rc *ReportController) GetByID(c echo.Context) error {
	report_id, _ := strconv.Atoi(c.Param("id"))

	report, err := rc.reportUseCase.GetByID(report_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportResponse := response_report.GetPaginateFromEntitiesToResponse(&report)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Report By ID", reportResponse))
}

func (rc *ReportController) Update(c echo.Context) error {
	user_id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	report_id, _ := strconv.Atoi(c.Param("id"))

	var reportRequest request.Update
	c.Bind(&reportRequest)
	reportRequest.UserID = user_id
	reportRequest.ID = report_id

	// Parse form-data multipart
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err.Error()))
	}

	report, err := rc.reportUseCase.Update(*reportRequest.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportResponse := response_report.UpdateFromEntitiesToResponse(&report)

	// Mengambil semua file yang diunggah
	files := form.File["files"]
	if len(files) != 0 {
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

		_, err2 := rc.reportFileUseCase.Delete(report_id)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err2.Error()))
		}

		reportFile, err3 := rc.reportFileUseCase.Create(files, report_id)
		if err3 != nil {
			return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err3.Error()))
		}

		reportFileResponses := []*response_report_file.ReportFile{}
		for _, rf := range reportFile {
			reportFileResponses = append(reportFileResponses, response_report_file.FromEntitiesToResponse(&rf))
		}

		reportResponse.Files = reportFileResponses
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Report", reportResponse))

}

func (rc *ReportController) Delete(c echo.Context) error {
	user_id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	report_id, _ := strconv.Atoi(c.Param("id"))

	report, err := rc.reportUseCase.Delete(report_id, user_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportResponse := response_report.DeleteFromEntitiesToResponse(&report)

	reportFile, err2 := rc.reportFileUseCase.Delete(report_id)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err2.Error()))
	}

	reportFileResponses := []*response_report_file.ReportFile{}
	for _, rf := range reportFile {
		reportFileResponses = append(reportFileResponses, response_report_file.FromEntitiesToResponse(&rf))
	}

	reportResponse.Files = reportFileResponses

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Report", reportResponse))
}

func (rc *ReportController) AdminDelete(c echo.Context) error {
	report_id, _ := strconv.Atoi(c.Param("id"))

	report, err := rc.reportUseCase.AdminDelete(report_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportResponse := response_report.DeleteFromEntitiesToResponse(&report)

	reportFile, err2 := rc.reportFileUseCase.Delete(report_id)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err2.Error()))
	}

	reportFileResponses := []*response_report_file.ReportFile{}
	for _, rf := range reportFile {
		reportFileResponses = append(reportFileResponses, response_report_file.FromEntitiesToResponse(&rf))
	}

	reportResponse.Files = reportFileResponses

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Report", reportResponse))
}

func (rc *ReportController) GetDistanceDuration(c echo.Context) error {
	report_id, _ := strconv.Atoi(c.Param("id"))
	admin_id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	distanceDuration, err := rc.reportUseCase.GetDistanceDuration(report_id, admin_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	distanceDurationResponse := response_report.DistanceDurationFromEntitiesToResponse(distanceDuration)

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Distance & Duration", distanceDurationResponse))

}
