package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	profileDaoObject dao.ProfileDao = dao.NewProfileDao()
)

func Profile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		profileRequest := ctx.Query("bcaId")
		currentTime := time.Now()
		var responseCode int
		var profileResponse response.ProfileResponse
		var outputSchema response.ProfileOutputSchema
		var errorSchema response.ErrorSchema
		if profileRequest == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := profileDaoObject.GetById(profileRequest)

			if result.BcaId == "" {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				responseCode = 404
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailProfile = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				profileResponse.OutputSchema = outputSchema
				responseCode = 200
			}
		}
		profileResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, profileResponse)
	}
}
