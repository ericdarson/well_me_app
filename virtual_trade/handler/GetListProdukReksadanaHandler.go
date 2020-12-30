package handler

import (
	"strings"
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

		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		b := strings.IndexFunc(idjenis, isNotDigit) == -1
		if !b {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID MUST ONLY CONTAIN A NUMBER"
			errorSchema.ErrorMessage.Indonesian = "ID HANYA BOLEH MENGANDUNG ANGKA"
			getListProdukReksadana.ErrorSchema = errorSchema
			ctx.JSON(400, getListProdukReksadana)
		} else {
			result := daoListProduk.GetListProdukReksadana(idjenis)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				getListProdukReksadana.ErrorSchema = errorSchema
				ctx.JSON(404, getListProdukReksadana)
			} else if result[0].Id == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				getListProdukReksadana.ErrorSchema = errorSchema
				ctx.JSON(500, getListProdukReksadana)
			} else {
				outputSchema = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				getListProdukReksadana.OutputSchema = outputSchema
				getListProdukReksadana.ErrorSchema = errorSchema
				ctx.JSON(200, getListProdukReksadana)
			}
		}
	}
}
