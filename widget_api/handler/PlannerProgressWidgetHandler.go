package handler

import (
	"fmt"
	"strings"
	"widget_api/dao"
	"widget_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoProgressPlanner dao.ProgressPlannerWidgetDao = dao.NewProgressPlanner()
)

func PlannerProgressWidgetGetByIds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		listIds := ctx.Query("IdsPlanner")
		fmt.Println("list id: " + listIds)
		var getListPlanner response.PlannerProgressWidgetResponse
		var outputSchema []response.PlannerProgressWidgetOutputSchema
		var errorSchema response.ErrorSchema

		isNotDigit := func(c rune) bool { return (c < '0' || c > '9') && c != ',' }
		b := strings.IndexFunc(listIds, isNotDigit) == -1
		if !b {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID MUST ONLY CONTAIN A NUMBER"
			errorSchema.ErrorMessage.Indonesian = "ID HANYA BOLEH MENGANDUNG ANGKA"
			getListPlanner.ErrorSchema = errorSchema
			ctx.JSON(400, getListPlanner)
		} else {
			result := daoProgressPlanner.GetProgressPlannerWidget(listIds)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				getListPlanner.ErrorSchema = errorSchema
				ctx.JSON(404, getListPlanner)
			} else if result[0].IdPlan == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				getListPlanner.ErrorSchema = errorSchema
				ctx.JSON(500, getListPlanner)
			} else {
				outputSchema = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				getListPlanner.OutputSchema = outputSchema
				getListPlanner.ErrorSchema = errorSchema
				ctx.JSON(200, getListPlanner)
			}
		}
	}
}
