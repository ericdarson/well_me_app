package handler

import (
	"widget_api/dao"
	"widget_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoMyPlannerList dao.MyPlannerListDao = dao.NewMyPlannerList()
)

func GetMyPlannerListWithFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bcaid := ctx.Query("bcaid")
		planName := ctx.Query("planName")
		sortBy := ctx.Query("sortBy")
		var getMyListPlanner response.GetMyPlannerListResponse
		var outputSchema []response.MyPlannerListOutputSchema
		var errorSchema response.ErrorSchema

		if bcaid == "" || sortBy == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			getMyListPlanner.ErrorSchema = errorSchema
			ctx.JSON(400, getMyListPlanner)
		} else {
			result := daoMyPlannerList.GetMyPlannerList(bcaid, planName, sortBy)
			if result == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				getMyListPlanner.ErrorSchema = errorSchema
				ctx.JSON(404, getMyListPlanner)
			} else if result[0].IdPlan == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				getMyListPlanner.ErrorSchema = errorSchema
				ctx.JSON(500, getMyListPlanner)
			} else {
				outputSchema = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				getMyListPlanner.OutputSchema = outputSchema
				getMyListPlanner.ErrorSchema = errorSchema
				ctx.JSON(200, getMyListPlanner)
			}
		}
	}
}
