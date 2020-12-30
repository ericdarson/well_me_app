package handler

import (
	"list_reksadana_api/dao"
	"list_reksadana_api/entity/response"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	daoListReksadana dao.ListReksadanaDao = dao.NewListReksadana()
)

func ListProdukReksadanaHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var listProdukReksadanaResponse response.ListProdukReksadanaResponse
		var errorSchema response.ErrorSchema

		idjenis := ctx.Param("id-jenis")

		if idjenis == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			listProdukReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, listProdukReksadanaResponse)
		} else {
			isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
			b := strings.IndexFunc(idjenis, isNotDigit) == -1
			if !b {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "ID CONTAIN NON-NUMERIC CHARACTER"
				errorSchema.ErrorMessage.Indonesian = "ID MENGANDUNG KARAKTER SELAIN ANGKA"
				listProdukReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(400, listProdukReksadanaResponse)
			} else {
				outputSchema := daoListReksadana.GetListProdukReksadana(idjenis)
				if outputSchema == nil {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
					listProdukReksadanaResponse.ErrorSchema = errorSchema
					ctx.JSON(404, listProdukReksadanaResponse)
				} else if outputSchema[0].ID == "-1" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					listProdukReksadanaResponse.ErrorSchema = errorSchema
					ctx.JSON(500, listProdukReksadanaResponse)
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					listProdukReksadanaResponse.OutputSchema = outputSchema
					listProdukReksadanaResponse.ErrorSchema = errorSchema
					ctx.JSON(200, listProdukReksadanaResponse)
				}
			}
		}
	}
}
