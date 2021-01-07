package handler

import (
	"planner_resiko_api/dao"
	"planner_resiko_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoPlannerResiko dao.PlannerResikoDao = dao.NewPlannerResiko()
)

func PlannerResikoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.PlannerResikoResponse
		var errorSchema response.ErrorSchema

		bcaid := ctx.Param("bca-id")

		if bcaid == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoPlannerResiko.GetResiko(bcaid)
			if outputSchema == nil {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if outputSchema[0].IDJenis == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.OutputSchema = outputSchema
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			}
		}
	}
}
