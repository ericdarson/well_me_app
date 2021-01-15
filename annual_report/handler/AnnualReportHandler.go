package handler

import (
	"annual_report/dao"
	"annual_report/entity/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	daoAnnualReport dao.AnnualReportDao = dao.NewAnnualReport()
)

func AnnualReportGetByBcaId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bcaId := ctx.Query("bcaid")
		fmt.Println("list id: " + bcaId)
		var getAnnualReport response.AnnualReportResponse
		var outputSchema []response.AnnualReportOutputSchema
		var errorSchema response.ErrorSchema

		//isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && c != ',' }
		//b := strings.IndexFunc(bcaid, isNotDigit) == -1
		if bcaId == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID MUST ONLY CONTAIN A NUMBER"
			errorSchema.ErrorMessage.Indonesian = "ID HANYA BOLEH MENGANDUNG ANGKA"
			getAnnualReport.ErrorSchema = errorSchema
			ctx.JSON(400, getAnnualReport)
		} else {
			result := daoAnnualReport.GetAnnualReport(bcaId)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				getAnnualReport.ErrorSchema = errorSchema
				ctx.JSON(404, getAnnualReport)
			} else if result[0].BcaId == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				getAnnualReport.ErrorSchema = errorSchema
				ctx.JSON(500, getAnnualReport)
			} else {
				outputSchema = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				getAnnualReport.OutputSchema = outputSchema
				getAnnualReport.ErrorSchema = errorSchema
				ctx.JSON(200, getAnnualReport)
			}
		}
	}
}
