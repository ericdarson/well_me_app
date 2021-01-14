package handler

import (
	"cms_api/dao"
	"cms_api/entity/request"
	"cms_api/entity/response"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	daoDaily dao.DailyNabCMSDao = dao.NewDailyNabCMS()
)

func GetDailyProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idjenis := ctx.Param("id-jenis")
		var JSONReponse response.DailyNabResponse
		var errorSchema response.ErrorSchema
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if idjenis == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "REQUIRED INPUT IS EMPTY"
			errorSchema.ErrorMessage.Indonesian = "INPUT YANG DIBUTUHKAN KOSONG"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(401, JSONReponse)
		} else if strings.IndexFunc(idjenis, isNotDigit) != -1 {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID JENIS SHOULD ONLY CONTAIN NUMBER"
			errorSchema.ErrorMessage.Indonesian = "ID JENIS HANYA DAPAT MENGANDUNG ANGKA"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(401, JSONReponse)
		} else {
			outputSchema := daoDaily.InquiryDailyNab(idjenis)
			if outputSchema == nil {
				errorSchema.ErrorCode = "BIT-17-002"
				errorSchema.ErrorMessage.English = "NO DATA FOUND"
				errorSchema.ErrorMessage.Indonesian = "TIDAK ADA DATA DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema[0].IDProduk == -1 {
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

func InsertDailyProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.InsertDailyProductsResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.InsertDailyProductsRequest
		err := ctx.BindJSON(&JSONRequest)
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "FORMAT REQUEST TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(401, JSONReponse)
		} else {
			outputSchema := daoDaily.InsertDailyNab(JSONRequest)
			if outputSchema.ListGagal == nil {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				JSONReponse.OutputSchema = outputSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema.ListGagal[0].IDProduk == -1 {
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
