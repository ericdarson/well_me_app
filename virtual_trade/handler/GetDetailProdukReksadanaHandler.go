package handler

import (
	"virtual_trade_api/dao"
	"virtual_trade_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoDetailProduk dao.GetDetailProdukReksadanaDao = dao.NewDetailProduk()
)

func GetDetailProdukReksadanaHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idproduk := ctx.Param("id-produk")
		timefilter := ctx.Param("filter")
		var detailProdukReksadanaResponse response.GetDetailProdukReksadanaResponse
		var outputSchema response.GetDetailProdukReksadanaOutputSchema
		var errorSchema response.ErrorSchema

		result := daoDetailProduk.GetDetailProdukReksadana(idproduk, timefilter)
		if result.ID == "-1" {
			errorSchema.ErrorCode = "BIT-17-005"
			errorSchema.ErrorMessage.English = "GENERAL ERROR"
			errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
		} else {
			outputSchema = result
			if outputSchema.ID == "" {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
			} else {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
			}
			detailProdukReksadanaResponse.OutputSchema = outputSchema
		}

		detailProdukReksadanaResponse.ErrorSchema = errorSchema
		ctx.JSON(200, detailProdukReksadanaResponse)
	}
}
