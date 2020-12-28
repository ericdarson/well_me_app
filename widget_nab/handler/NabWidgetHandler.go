package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoObject dao.NabWidgetDao = dao.New()
)

func NabWidgetGetByIds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		listIds := ctx.Query("ids")

		currentTime := time.Now()
		var nabWidgetResponse response.NabWidgetResponse
		var outputSchema response.NabWidgetOutputSchema
		var errorSchema response.ErrorSchema
		if listIds == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
		} else {
			result := daoObject.GetByIds(listIds)
			if result[0].Reksadana == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.ListNAB = result
				if outputSchema.ListNAB == nil {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				}
				nabWidgetResponse.OutputSchema = outputSchema
			}
		}
		nabWidgetResponse.ErrorSchema = errorSchema
		ctx.JSON(200, nabWidgetResponse)
	}
}
