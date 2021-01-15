package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	loginAdminDaoObject dao.LoginAdminDao = dao.NewLoginAdminDao()
)

type LoginAdminAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		responseCode := 200
		loginAdminRequest := LoginAdminAdminRequest{}
		ctx.Bind(&loginAdminRequest)
		currentTime := time.Now()
		var loginAdminResponse response.LoginAdminResponse
		var outputSchema response.LoginAdminOutputSchema
		var errorSchema response.ErrorSchema
		if loginAdminRequest.Username == "" || loginAdminRequest.Password == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			responseCode = 400
		} else {
			result := loginAdminDaoObject.LoginAdmin(ctx, loginAdminRequest.Username, loginAdminRequest.Password)
			if result.Message == "" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				responseCode = 500
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailLoginAdmin = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				loginAdminResponse.OutputSchema = outputSchema
				responseCode = 200
			}
		}
		loginAdminResponse.ErrorSchema = errorSchema
		ctx.JSON(responseCode, loginAdminResponse)
	}
}
