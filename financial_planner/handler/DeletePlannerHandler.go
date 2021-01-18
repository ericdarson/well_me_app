package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	deletePlannerDaoObject dao.DeletePlannerDao = dao.NewDeletePlannerDao()
)

type DeletePlannerRequest struct {
	BcaId  string `json:"bcaId"`
	IdPlan string `json:"idPlan"`
}

func DeletePlanner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		deletePlannerRequest := DeletePlannerRequest{}
		ctx.Bind(&deletePlannerRequest)
		currentTime := time.Now()
		var deletePlannerResponse response.DeletePlannerResponse
		var outputSchema response.DeletePlannerOutputSchema
		var errorSchema response.ErrorSchema
		if deletePlannerRequest.BcaId == "" || deletePlannerRequest.IdPlan == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := deletePlannerDaoObject.DeletePlanner(deletePlannerRequest.IdPlan, deletePlannerRequest.BcaId)
			if result.Message == "0" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				responseCode = 500
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailDeletePlanner = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"

				deletePlannerResponse.OutputSchema = outputSchema
			}
		}
		deletePlannerResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, deletePlannerResponse)
	}
}
