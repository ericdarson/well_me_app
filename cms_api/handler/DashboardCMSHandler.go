package handler

import (
	"cms_api/dao"
	"cms_api/entity/response"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	daoDashboard dao.DashboardCMSDao = dao.NewDashboardCMS()
)

func InquiryDashboardOverview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

		startDate := ctx.Param("start-date")
		endDate := ctx.Param("end-date")
		chartType := ctx.Param("chart-type")

		var JSONReponse response.DashboardOverviewResponse
		var errorSchema response.ErrorSchema
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if !(re.MatchString(strings.ReplaceAll(startDate, "-", "/"))) || !(re.MatchString(strings.ReplaceAll(endDate, "-", "/"))) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if chartType != "daily" && chartType != "weekly" && chartType != "monthly" && chartType != "yearly" {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID CHART TYPE"
			errorSchema.ErrorMessage.Indonesian = "TIPE GRAFIK TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			if chartType == "daily" {
				chartType = "dd-mm-yyyy"
			} else if chartType == "weekly" {
				chartType = "IW"
			} else if chartType == "monthly" {
				chartType = "mm-yyyy"
			} else if chartType == "yearly" {
				chartType = "yyyy"
			}
			outputSchema := daoDashboard.InquiryOverview(chartType, startDate, endDate)
			if outputSchema.User == -1 {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema.User == -2 {
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

func InquiryPromoUsage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.DashboardPromoResponse
		var errorSchema response.ErrorSchema
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else {
			outputSchema := daoDashboard.InquiryPromoUsage()
			if outputSchema.Objectives == nil && outputSchema.Promotions == nil {
				errorSchema.ErrorCode = "BIT-17-003"
				errorSchema.ErrorMessage.English = "NO DATA FOUND"
				errorSchema.ErrorMessage.Indonesian = "TIDAK ADA DATA"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema.Objectives[0].KodePromo == "-1" {
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
