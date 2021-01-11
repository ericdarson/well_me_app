package handler

import (
	"fmt"
	"strings"
	"widget_api/dao"
	"widget_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoDailyProfit dao.DailyProfitWidgetDao = dao.NewDailyProfit()
)

func DailyProfitWidgetGetByIds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		listIds := ctx.Query("listid")
		fmt.Println("list id: " + listIds)
		var getListDailyProfit response.GetDailyProfitWidgetResponse
		var outputSchema []response.GetDailyProfitWidgetOutputSchema
		var errorSchema response.ErrorSchema

		isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && c != ',' }
		b := strings.IndexFunc(listIds, isNotDigit) == -1
		if !b {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID MUST ONLY CONTAIN A NUMBER"
			errorSchema.ErrorMessage.Indonesian = "ID HANYA BOLEH MENGANDUNG ANGKA"
			getListDailyProfit.ErrorSchema = errorSchema
			ctx.JSON(400, getListDailyProfit)
		} else {
			result := daoDailyProfit.GetDailyProfitWidget(listIds)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				getListDailyProfit.ErrorSchema = errorSchema
				ctx.JSON(404, getListDailyProfit)
			} else if result[0].Id == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				getListDailyProfit.ErrorSchema = errorSchema
				ctx.JSON(500, getListDailyProfit)
			} else {
				outputSchema = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				getListDailyProfit.OutputSchema = outputSchema
				getListDailyProfit.ErrorSchema = errorSchema
				ctx.JSON(200, getListDailyProfit)
			}
		}
	}
}
