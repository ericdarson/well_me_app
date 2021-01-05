package handler

import (
	"strings"
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
		isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && c != ',' }
		b := strings.IndexFunc(listIds, isNotDigit) == -1
		if listIds == "" || strings.Contains(listIds, ",,") {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			nabWidgetResponse.ErrorSchema = errorSchema
			ctx.JSON(400, nabWidgetResponse)
		} else if !b {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID CAN ONLY CONTAINS NUMBER AND IS SEPARATED BY COMMA"
			errorSchema.ErrorMessage.Indonesian = "ID HANYA BOLEH BERISI ANGKA DAN DIPISAHKAN DENGAN TANDA BACA KOMA"
			nabWidgetResponse.ErrorSchema = errorSchema
			ctx.JSON(400, nabWidgetResponse)
		} else {
			result := daoObject.GetByIds(listIds)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				nabWidgetResponse.ErrorSchema = errorSchema
				ctx.JSON(404, nabWidgetResponse)
			} else if result[0].Reksadana == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
        nabWidgetResponse.ErrorSchema = errorSchema
				ctx.JSON(500, nabWidgetResponse)
			} else {
				outputSchema.SystemDate = currentTime.Format("02-01-2006")
				outputSchema.ListNAB = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				nabWidgetResponse.OutputSchema = outputSchema
				nabWidgetResponse.ErrorSchema = errorSchema
				ctx.JSON(200, nabWidgetResponse)
			}
		}
	}
}
