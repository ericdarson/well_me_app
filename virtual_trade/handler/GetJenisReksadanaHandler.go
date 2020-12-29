package handler

import (
	"virtual_trade_api/dao"
	"virtual_trade_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoObject dao.GetJenisReksanaDanaDao = dao.New()
)

func GetJenisReksadanaHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var getJenisReksadanaResponse response.GetJenisReksadanaResponse
		var outputSchema []response.GetJenisReksadanaOutputSchema
		var errorSchema response.ErrorSchema

		result := daoObject.GetJenisReksadana()
		if result[0].Id == "-1" {
			errorSchema.ErrorCode = "BIT-17-005"
			errorSchema.ErrorMessage.English = "GENERAL ERROR"
			errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
		} else {
			outputSchema = result
			if outputSchema == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
			} else {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
			}
			getJenisReksadanaResponse.OutputSchema = outputSchema
		}

		getJenisReksadanaResponse.ErrorSchema = errorSchema
		ctx.JSON(200, getJenisReksadanaResponse)
	}
}
