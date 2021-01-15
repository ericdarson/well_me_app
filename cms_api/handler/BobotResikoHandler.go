package handler

import (
	"cms_api/dao"
	"cms_api/entity/request"
	"cms_api/entity/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	daoBobotResiko dao.BobotResikoCMSDao = dao.NewBobotResikoCMS()
)

func InquiryBobotResiko() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.InquiryBobotResikoResponse
		var errorSchema response.ErrorSchema

		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else {
			outputSchema := daoBobotResiko.InquiryBobotResiko()

			if outputSchema == nil {
				errorSchema.ErrorCode = "BIT-17-003"
				errorSchema.ErrorMessage.English = "NO DATA FOUND"
				errorSchema.ErrorMessage.Indonesian = "TIDAK ADA DATA YANG DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema[0].BobotResiko == "-1" {
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

func UpdateBobotResiko() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var JSONReponse response.OnlyErrorSchemaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.UpdateBobotResikoRequest
		err := ctx.Bind(&JSONRequest)

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
			ctx.JSON(403, JSONReponse)
		} else if !validateUpdateBobot(JSONRequest) {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "TOTAL PERCENTAGE IS NOT 100"
			errorSchema.ErrorMessage.Indonesian = "TOTAL PERSENTASE TIDAK 100"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else {
			outputSchema := daoBobotResiko.UpdateBobotResiko(JSONRequest)
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
			} else if outputSchema == "NOT FOUND" {
				errorSchema.ErrorCode = "BIT-17-003"
				errorSchema.ErrorMessage.English = "NO DATA FOUND"
				errorSchema.ErrorMessage.Indonesian = "TIDAK ADA DATA YANG DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
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

func validateUpdateBobot(req request.UpdateBobotResikoRequest) bool {
	m := make(map[float64]float64)
	for _, single := range req.Input {
		m[single.BobotResiko] += single.Persentase
	}

	for _, percentage := range m {
		if percentage != 100 {
			return false
		}
	}
	return true
}
