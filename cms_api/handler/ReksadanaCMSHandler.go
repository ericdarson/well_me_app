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
	daoReksadana dao.ReksadanaDao = dao.NewReksadana()
)

func InquiryAllJenisReksadana() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.JenisReksadanaResponse
		var errorSchema response.ErrorSchema

		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else {
			outputSchema := daoReksadana.InquiryAllJenisReksadana()
			if outputSchema == nil {
				errorSchema.ErrorCode = "BIT-17-002"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema[0].ID == "-1" {
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

func InsertJenisReksadana() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.InsertJenisReksadanaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.InsertJenisReksadanaRequest

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
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoReksadana.InsertJenisReksadana(JSONRequest)
			if outputSchema.ID == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema.ID == "-2" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
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

func UpdateJenisReksadana() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.InsertJenisReksadanaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.InsertJenisReksadanaRequest

		err := ctx.BindJSON(&JSONRequest)
		idJenis := ctx.Param("id-jenis")
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
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
			ctx.JSON(400, JSONReponse)
		} else if idJenis == "" || strings.IndexFunc(idJenis, isNotDigit) != -1 {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID JENIS INVALID"
			errorSchema.ErrorMessage.Indonesian = "FORMAT ID JENIS TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoReksadana.UpdateJenisReksadana(JSONRequest, idJenis)
			if outputSchema.ID == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema.ID == "-2" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
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

func InquiryProdukReksadana() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.InquiryProdukReksadanaResponse
		var errorSchema response.ErrorSchema

		idJenis := ctx.Query("id-jenis")
		filter := "IN"
		if idJenis == "" {
			idJenis = "0"
			filter = "NOT IN"
		}

		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else {
			outputSchema := daoReksadana.InquiryProdukReksadana(idJenis, filter)
			if outputSchema == nil {
				errorSchema.ErrorCode = "BIT-17-003"
				errorSchema.ErrorMessage.English = "NO DATA FOUND"
				errorSchema.ErrorMessage.Indonesian = "TIDAK ADA DATA YANG DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema[0].ID == "-1" {
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

func InsertProdukReksadana() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.InsertProdukReksadanaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.ProdukReksadanaRequest

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
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoReksadana.InsertProdukReksadana(JSONRequest)
			if outputSchema == "ERROR" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				JSONReponse.OutputSchema = response.InsertProdukReksadanaOutputSchema{ID: outputSchema}
				ctx.JSON(200, JSONReponse)
			}
		}
	}
}

func UpdateProdukReksadana() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.OnlyErrorSchemaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.ProdukReksadanaRequest

		err := ctx.BindJSON(&JSONRequest)
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		idProduk := ctx.Param("id-produk")
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
			ctx.JSON(400, JSONReponse)
		} else if idProduk == "" || strings.IndexFunc(idProduk, isNotDigit) != -1 {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID ID PRODUK"
			errorSchema.ErrorMessage.Indonesian = "FORMAT ID PRODUK TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoReksadana.UpdateProdukReksadana(JSONRequest, idProduk)
			if outputSchema == "ERROR" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			}
		}
	}
}
