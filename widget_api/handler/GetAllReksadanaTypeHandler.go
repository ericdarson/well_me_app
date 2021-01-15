package handler

import (
	"widget_api/dao"
	"widget_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoObj dao.AllReksadanaTypeDao = dao.NewAllReksadanaType()
)

func GetAllReksadanaType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var allReksadanaTypeResponse response.GetAllReksadanaTypeResponse
		var outputSchema []response.GetAllReksadanaTypeOutputSchema
		var errorSchema response.ErrorSchema

		result := daoObj.GetAllType()
		if result == nil {
			errorSchema.ErrorCode = "BIT-17-004"
			errorSchema.ErrorMessage.English = "DATA NOT FOUND"
			errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
			allReksadanaTypeResponse.ErrorSchema = errorSchema
			ctx.JSON(404, allReksadanaTypeResponse)
		} else if result[0].Id == "-1" {
			errorSchema.ErrorCode = "BIT-17-005"
			errorSchema.ErrorMessage.English = "GENERAL ERROR"
			errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
			allReksadanaTypeResponse.ErrorSchema = errorSchema
			ctx.JSON(500, allReksadanaTypeResponse)
		} else {
			outputSchema = result
			errorSchema.ErrorCode = "BIT-00-000"
			errorSchema.ErrorMessage.English = "SUCCESS"
			errorSchema.ErrorMessage.Indonesian = "BERHASIL"
			allReksadanaTypeResponse.OutputSchema = outputSchema
			allReksadanaTypeResponse.ErrorSchema = errorSchema
			ctx.JSON(200, allReksadanaTypeResponse)
		}
	}
}
