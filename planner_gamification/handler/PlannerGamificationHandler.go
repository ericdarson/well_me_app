package handler

import (
	"planner_gamification_api/dao"
	"planner_gamification_api/entity/response"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	daoPlannerGamification dao.PlannerGamificationDao = dao.NewPlannerGamification()
)

func PlannerGamificationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.PlannerGamificationResponse
		var errorSchema response.ErrorSchema

		idplan := ctx.Param("id-plan")

		if idplan == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
			b := strings.IndexFunc(idplan, isNotDigit) == -1
			if !b {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "ID CONTAIN NON-NUMERIC CHARACTER"
				errorSchema.ErrorMessage.Indonesian = "ID MENGANDUNG KARAKTER SELAIN ANGKA"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else {
				outputSchema := daoPlannerGamification.GetDetailPlan(idplan)
				if outputSchema.Nama == "" {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(404, JSONReponse)
				} else if outputSchema.Nama == "-1" {
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
}
