package report_upvote

import (
	"lawan-tambang-liar/controllers/base"
	"lawan-tambang-liar/controllers/report_upvote/response"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReportUpvoteController struct {
	reportUseCase       entities.ReportUseCaseInterface
	reportUpvoteUseCase entities.ReportUpvoteUseCaseInterface
}

func NewReportUpvoteController(reportUseCase entities.ReportUseCaseInterface, reportUpvoteUseCase entities.ReportUpvoteUseCaseInterface) *ReportUpvoteController {
	return &ReportUpvoteController{
		reportUseCase:       reportUseCase,
		reportUpvoteUseCase: reportUpvoteUseCase,
	}
}

func (rc *ReportUpvoteController) ToggleUpvote(c echo.Context) error {
	userID, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	reportID, _ := strconv.Atoi(c.Param("id"))

	reportUpvote, status, err := rc.reportUpvoteUseCase.ToggleUpvote(userID, reportID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	if status == "upvoted" {
		err = rc.reportUseCase.IncreaseUpvote(reportID)
		if err != nil {
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, base.NewSuccessResponse("Success upvoted", response.FromEntitiesToResponse(&reportUpvote)))
	} else {
		err = rc.reportUseCase.DecreaseUpvote(reportID)
		if err != nil {
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, base.NewSuccessResponse("Success cancel upvoted", response.FromEntitiesToResponse(&reportUpvote)))
	}
}
