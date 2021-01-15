package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	checkSessionAdminDaoObject dao.CheckSessionAdminDao = dao.NewCheckSessionAdminDao()
)

func CheckSessionAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		token := ctx.Query("token")
		username := ctx.Query("username")
		currentTime := time.Now()
		var checkSessionAdminResponse response.CheckSessionAdminResponse
		var outputSchema response.CheckSessionAdminOutputSchema
		var errorSchema response.ErrorSchema
		if token == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := checkSessionAdminDaoObject.CheckSessionAdmin(username, token)
			if result.Message == "Gagal" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				responseCode = 500
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailCheckSession = result
				if outputSchema.DetailCheckSession.Message == "DATA TIDAK DITEMUKAN" {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
					responseCode = 200
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				}
				checkSessionAdminResponse.OutputSchema = outputSchema
			}
		}
		checkSessionAdminResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, checkSessionAdminResponse)
	}
}
