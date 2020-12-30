package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	logoutDaoObject dao.LogoutDao = dao.NewLogoutDao()
)

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		logoutRequest := ctx.Query("token")
		currentTime := time.Now()
		var logoutResponse response.LogoutResponse
		var outputSchema response.LogoutOutputSchema
		var errorSchema response.ErrorSchema
		if logoutRequest == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := logoutDaoObject.Logout(logoutRequest)
			if result.Message == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				responseCode = 500
			} else if result.Message == "DATA TIDAK DITEMUKAN" {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailLogout = result
				logoutResponse.OutputSchema = outputSchema
				responseCode = 404
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailLogout = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				logoutResponse.OutputSchema = outputSchema
				responseCode = 200
			}
		}
		logoutResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, logoutResponse)
	}
}
