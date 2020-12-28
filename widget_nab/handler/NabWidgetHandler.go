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
			errorSchema.ErrorCode = "BIT-77-777"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
		} else {
			result := daoObject.GetByIds(listIds)
			if result[0].Reksadana == "-1" {
				errorSchema.ErrorCode = "BIT-88-888"
				errorSchema.ErrorMessage.English = "SYSTEM MAINTENANCE"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.ListNAB = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				nabWidgetResponse.OutputSchema = outputSchema
			}
		}
		nabWidgetResponse.ErrorSchema = errorSchema
		ctx.JSON(200, nabWidgetResponse)
	}
}
