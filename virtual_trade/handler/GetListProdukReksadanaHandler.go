package handler

import (
	"virtual_trade_api/dao"
	"virtual_trade_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoListProduk dao.GetListProdukReksadanaDao = dao.NewListProduk()
)

func GetListProdukReksadanaHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idjenis := ctx.Param("id-jenis")
		var getListProdukReksadana response.GetListProdukReksadanaResponse
		var outputSchema []response.GetListProdukReksadanaOutputSchema
		var errorSchema response.ErrorSchema

		result := daoListProduk.GetListProdukReksadana(idjenis)
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
			getListProdukReksadana.OutputSchema = outputSchema
		}

		getListProdukReksadana.ErrorSchema = errorSchema
		ctx.JSON(200, getListProdukReksadana)
	}
}
