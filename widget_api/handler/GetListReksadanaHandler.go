package handler

import (
	"strings"
	"widget_api/dao"
	"widget_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoListReksadana   dao.ListReksadanaWidgetDao   = dao.NewListReksadana()
	daoMyListReksadana dao.MyListReksadanaWidgetDao = dao.NewMyListReksadana()
)

func GetListReksadanaWithFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nama := ctx.Query("nama")
		listKategori := ctx.Query("kategori")
		sortBy := ctx.Query("sortBy")

		var listReksadanaResponse response.ListReksadanaWidgetResponse
		var outputSchema []response.ListReksadanaWidgetOutputSchema
		var errorSchema response.ErrorSchema
		isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && c != ',' }
		b := strings.IndexFunc(listKategori, isNotDigit) == -1
		if listKategori == "" || strings.Contains(listKategori, ",,") {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			listReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, listReksadanaResponse)
		} else if !b {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID KATEGORI CAN ONLY CONTAINS NUMBER AND IS SEPARATED BY COMMA"
			errorSchema.ErrorMessage.Indonesian = "ID KATEGORI HANYA BOLEH BERISI ANGKA DAN DIPISAHKAN DENGAN TANDA BACA KOMA"
			listReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, listReksadanaResponse)
		} else {
			result := daoListReksadana.GetReksadanaList(nama, listKategori, sortBy)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				listReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(404, listReksadanaResponse)
			} else if result[0].Id == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				listReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(500, listReksadanaResponse)
			} else {
				outputSchema = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				listReksadanaResponse.OutputSchema = outputSchema
				listReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(200, listReksadanaResponse)
			}
		}
	}
}

func GetMyListReksadanaWithFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bcaid := ctx.Query("bcaid")
		nama := ctx.Query("nama")
		listKategori := ctx.Query("kategori")
		sortBy := ctx.Query("sortBy")

		var listReksadanaResponse response.ListReksadanaWidgetResponse
		var outputSchema []response.ListReksadanaWidgetOutputSchema
		var errorSchema response.ErrorSchema
		isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && c != ',' }
		b := strings.IndexFunc(listKategori, isNotDigit) == -1
		if listKategori == "" || strings.Contains(listKategori, ",,") || bcaid == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			listReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, listReksadanaResponse)
		} else if !b {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID KATEGORI CAN ONLY CONTAINS NUMBER AND IS SEPARATED BY COMMA"
			errorSchema.ErrorMessage.Indonesian = "ID KATEGORI HANYA BOLEH BERISI ANGKA DAN DIPISAHKAN DENGAN TANDA BACA KOMA"
			listReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, listReksadanaResponse)
		} else {
			result := daoMyListReksadana.GetMyReksadanaList(bcaid, nama, listKategori, sortBy)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				listReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(404, listReksadanaResponse)
			} else if result[0].Id == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				listReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(500, listReksadanaResponse)
			} else {
				outputSchema = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				listReksadanaResponse.OutputSchema = outputSchema
				listReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(200, listReksadanaResponse)
			}
		}
	}
}
