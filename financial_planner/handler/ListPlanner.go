package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	listPlannerDao dao.ListPlannerDao = dao.NewListPlannerDao()
)

func ListPlanner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		bcaId := ctx.Query("bcaId")

		currentTime := time.Now()
		var nabWidgetResponse response.ListPlannerResponse
		var outputSchema response.ListPlannerOutputSchema
		var errorSchema response.ErrorSchema
		if bcaId == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := listPlannerDao.GetById(bcaId)
			if result[0].IdPlan == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				responseCode = 500
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.ListPlanner = result
				if outputSchema.ListPlanner[0].IdPlan == "-2" {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
					responseCode = 404
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					responseCode = 200
				}
				nabWidgetResponse.OutputSchema = outputSchema
			}
		}
		nabWidgetResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, nabWidgetResponse)
	}
}
