package handler

import (
	"cms_api/dao"
	"cms_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoNasabah dao.NasabahCMSDao = dao.NewNasabahCMS()
)

func GetDataNasabah() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bcaid := ctx.Param("bca-id")
		if bcaid == "" {
			bcaid = "%"
		}

		var JSONReponse response.InquiryDataNasabahResponse
		var errorSchema response.ErrorSchema

		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else {
			outputSchema := daoNasabah.InquiryDataNasabah(bcaid)
			if outputSchema.Nama == "" {
				errorSchema.ErrorCode = "BIT-17-002"
				errorSchema.ErrorMessage.English = "NO DATA FOUND"
				errorSchema.ErrorMessage.Indonesian = "TIDAK ADA DATA DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema.Nama == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				JSONReponse.OutputSchema = outputSchema
				ctx.JSON(200, JSONReponse)
			}
		}
	}
}
