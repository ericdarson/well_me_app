package handler

import (
	"fmt"
	"strconv"
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	targetSimulationDao dao.TargetSimulationPlannerDao = dao.NewTargetSimulationPlanner()
)

type TargetSimulationRequest struct {
	BcaId    string `json:"bcaId"`
	Target   string `json:"target"`
	DueDate  string `json:"dueDate"`
	Periodic string `json:"periodic"`
}

func TargetSimulationPlanner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		targetSimulationRequest := TargetSimulationRequest{}

		ctx.Bind(&targetSimulationRequest)

		currentTime := time.Now()
		var targetSimulationResponse response.TargetSimulationResponse
		var outputSchema response.TargetSimulationOutputSchema
		var errorSchema response.ErrorSchema
		target, errfloat := strconv.ParseFloat(targetSimulationRequest.Target, 64)
		layout := "02-01-2006"
		_, errDate := time.Parse(layout, targetSimulationRequest.DueDate)
		if errfloat != nil || errDate != nil || (targetSimulationRequest.Periodic != "Weekly" && targetSimulationRequest.Periodic != "Monthly" && targetSimulationRequest.Periodic != "Yearly") {
			fmt.Println("errDate: ", errDate, "errfloat:", errfloat, "periodic: ", targetSimulationRequest.Periodic)
			responseCode := 400
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			targetSimulationResponse.ErrorSchema = errorSchema
			ctx.JSON(responseCode, targetSimulationResponse)
		} else {
			result, baseperiodic := targetSimulationDao.Simulation(targetSimulationRequest.BcaId, target, targetSimulationRequest.DueDate, targetSimulationRequest.Periodic)
			if result[0].NominalInvestasi == "-1" {
				responseCode = 500
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				targetSimulationResponse.ErrorSchema = errorSchema
				ctx.JSON(responseCode, targetSimulationResponse)
			} else {
				outputSchema.SystemDate = currentTime.Format("02-01-2006")
				outputSchema.TargetSimulationDetail = result
				outputSchema.TabunganPeriode = baseperiodic
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				targetSimulationResponse.OutputSchema = outputSchema
				targetSimulationResponse.ErrorSchema = errorSchema
				ctx.JSON(responseCode, targetSimulationResponse)
			}
		}
	}
}
