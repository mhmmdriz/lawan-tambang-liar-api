package report_solution_process

import (
	"lawan-tambang-liar/constants"
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/controllers/report_solution_process/request"
	response_report_solution "lawan-tambang-liar/controllers/report_solution_process/response"
	response_report_solution_file "lawan-tambang-liar/controllers/report_solution_process_file/response"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReportSolutionProcessController struct {
	reportUseCase             entities.ReportUseCaseInterface
	reportSolutionUseCase     entities.ReportSolutionProcessUseCaseInterface
	reportSolutionFileUseCase entities.ReportSolutionProcessFileUseCaseInterface
}

func NewReportSolutionProcessController(reportUseCase entities.ReportUseCaseInterface, reportSolutionUseCase entities.ReportSolutionProcessUseCaseInterface, reportSolutionFileUseCase entities.ReportSolutionProcessFileUseCaseInterface) *ReportSolutionProcessController {
	return &ReportSolutionProcessController{
		reportUseCase:             reportUseCase,
		reportSolutionUseCase:     reportSolutionUseCase,
		reportSolutionFileUseCase: reportSolutionFileUseCase,
	}
}

func (rc *ReportSolutionProcessController) Create(c echo.Context) error {
	admin_id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	var reportSolutionRequest request.Create
	c.Bind(&reportSolutionRequest)

	reportSolutionRequest.AdminID = admin_id
	report_id, _ := strconv.Atoi(c.Param("id"))
	reportSolutionRequest.ReportID = report_id
	action := c.Param("action")
	if action == "verify" {
		reportSolutionRequest.Status = "verified"
	} else if action == "reject" {
		reportSolutionRequest.Status = "rejected"
	} else if action == "progress" {
		reportSolutionRequest.Status = "on progress"
	} else if action == "finish" {
		reportSolutionRequest.Status = "done"
	} else {
		return c.JSON(utils.ConvertResponseCode(constants.ErrActionNotFound), base.NewErrorResponse(constants.ErrActionNotFound.Error()))
	}

	// Parse form-data multipart
	form, err2 := c.MultipartForm()
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err2.Error()))
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

	err3 := rc.reportUseCase.UpdateStatus(report_id, reportSolutionRequest.Status)
	if err3 != nil {
		return c.JSON(utils.ConvertResponseCode(err3), base.NewErrorResponse(err3.Error()))
	}

	reportSolution, err4 := rc.reportSolutionUseCase.Create(reportSolutionRequest.ToEntities())
	if err4 != nil {
		return c.JSON(utils.ConvertResponseCode(err4), base.NewErrorResponse(err4.Error()))
	}

	reportSolutionResponse := response_report_solution.CreateFromEntitiesToResponse(&reportSolution)

	reportSolutionFile, err5 := rc.reportSolutionFileUseCase.Create(files, reportSolution.ID)
	if err5 != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err5.Error()))
	}

	reportSolutionFileResponses := []*response_report_solution_file.ReportSolutionProcessFile{}
	for _, rf := range reportSolutionFile {
		reportSolutionFileResponses = append(reportSolutionFileResponses, response_report_solution_file.FromEntitiesToResponse(&rf))
	}

	reportSolutionResponse.Files = reportSolutionFileResponses

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Report Solution Process", reportSolutionResponse))
}

func (rc *ReportSolutionProcessController) GetByReportID(c echo.Context) error {
	report_id, _ := strconv.Atoi(c.Param("id"))

	reportSolutionProcesses, err := rc.reportSolutionUseCase.GetByReportID(report_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportSolutionProcessResponses := []*response_report_solution.GetByReportID{}
	for _, rsp := range reportSolutionProcesses {
		reportSolutionProcessResponses = append(reportSolutionProcessResponses, response_report_solution.GetByReportIDFromEntitiesToResponse(&rsp))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Report Solution Process", reportSolutionProcessResponses))
}

func (rc *ReportSolutionProcessController) Delete(c echo.Context) error {
	report_id, _ := strconv.Atoi(c.Param("id"))
	report_solution_id, _ := strconv.Atoi(c.Param("solution_id"))

	reportSolutionProcess, err := rc.reportSolutionUseCase.Delete(report_solution_id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportSolutionProcessResponse := response_report_solution.DeleteFromEntitiesToResponse(&reportSolutionProcess)

	var updatedStatus string
	if reportSolutionProcess.Status == "verified" {
		updatedStatus = "pending"
	} else if reportSolutionProcess.Status == "rejected" {
		updatedStatus = "pending"
	} else if reportSolutionProcess.Status == "on progress" {
		updatedStatus = "verified"
	} else if reportSolutionProcess.Status == "done" {
		updatedStatus = "on progress"
	}

	err2 := rc.reportUseCase.UpdateStatus(report_id, updatedStatus)
	if err2 != nil {
		return c.JSON(utils.ConvertResponseCode(err2), base.NewErrorResponse(err2.Error()))
	}

	reportSolutionFile, err3 := rc.reportSolutionFileUseCase.Delete(report_solution_id)
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err3.Error()))
	}

	reportSolutionFileResponses := []*response_report_solution_file.ReportSolutionProcessFile{}
	for _, rf := range reportSolutionFile {
		reportSolutionFileResponses = append(reportSolutionFileResponses, response_report_solution_file.FromEntitiesToResponse(&rf))
	}

	reportSolutionProcessResponse.Files = reportSolutionFileResponses

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Report Solution Process", reportSolutionProcessResponse))
}

func (rc *ReportSolutionProcessController) Update(c echo.Context) error {
	var reportSolutionRequest request.Update
	reportSolutionRequest.Message = c.FormValue("message")

	admin_id, err1 := utils.GetUserIDFromJWT(c)
	if err1 != nil {
		return c.JSON(utils.ConvertResponseCode(err1), base.NewErrorResponse(err1.Error()))
	}
	reportSolutionRequest.AdminID = admin_id
	report_id, _ := strconv.Atoi(c.Param("id"))
	reportSolutionRequest.ReportID = report_id
	report_solution_id, _ := strconv.Atoi(c.Param("solution_id"))
	reportSolutionRequest.ID = report_solution_id

	reportSolution, err2 := rc.reportSolutionUseCase.Update(*reportSolutionRequest.ToEntities())
	if err2 != nil {
		return c.JSON(utils.ConvertResponseCode(err2), base.NewErrorResponse(err2.Error()))
	}

	reportSolutionResponse := response_report_solution.UpdateFromEntitiesToResponse(&reportSolution)

	// Parse form-data multipart
	form, err2 := c.MultipartForm()
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err2.Error()))
	}

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

		_, err3 := rc.reportSolutionFileUseCase.Delete(report_solution_id)
		if err3 != nil {
			return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err3.Error()))
		}

		reportSolutionFile, err4 := rc.reportSolutionFileUseCase.Create(files, report_solution_id)
		if err4 != nil {
			return c.JSON(http.StatusInternalServerError, base.NewErrorResponse(err4.Error()))
		}

		reportSolutionFileResponses := []*response_report_solution_file.ReportSolutionProcessFile{}
		for _, rf := range reportSolutionFile {
			reportSolutionFileResponses = append(reportSolutionFileResponses, response_report_solution_file.FromEntitiesToResponse(&rf))
		}

		reportSolutionResponse.Files = reportSolutionFileResponses
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Report Solution Process", reportSolutionResponse))
}
