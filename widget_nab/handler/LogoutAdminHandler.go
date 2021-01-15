package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	logoutAdminDaoObject dao.LogoutAdminDao = dao.NewLogoutAdminDao()
)

func LogoutAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseCode := 200
		logoutAdminRequest := ctx.Request.Header["Token"]
		currentTime := time.Now()
		var logoutAdminResponse response.LogoutResponse
		var outputSchema response.LogoutOutputSchema
		var errorSchema response.ErrorSchema
		if logoutAdminRequest != nil {
			if logoutAdminRequest[0] == "" {
				errorSchema.ErrorCode = "BIT-17-002"
				errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
				errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
				responseCode = 400
			} else {
				result := logoutAdminDaoObject.LogoutAdmin(logoutAdminRequest[0])
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
					logoutAdminResponse.OutputSchema = outputSchema
					responseCode = 404
				} else {
					outputSchema.SystemDate = currentTime.Format("2006-01-02")
					outputSchema.DetailLogout = result
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					logoutAdminResponse.OutputSchema = outputSchema
					responseCode = 200
				}
			}
			logoutAdminResponse.ErrorSchema = errorSchema
			ctx.JSON(responseCode, logoutAdminResponse)
		} else {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			ctx.JSON(400, logoutAdminResponse)
		}
	}

}
