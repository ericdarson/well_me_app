package handler

import (
	"time"
	"widget_nab_service/dao"
	"widget_nab_service/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	loginDaoObject dao.LoginDao = dao.NewLoginDao()
)

type LoginRequest struct {
	BcaId    string `json:"bcaId"`
	Password string `json:"password"`
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginRequest := LoginRequest{}
		ctx.Bind(&loginRequest)
		currentTime := time.Now()
		var loginResponse response.LoginResponse
		var outputSchema response.LoginOutputSchema
		var errorSchema response.ErrorSchema
		if loginRequest.BcaId == "" || loginRequest.Password == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
		} else {
			result := loginDaoObject.Login(ctx, loginRequest.BcaId, loginRequest.Password)
			if result.Message == "" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
			} else {
				outputSchema.SystemDate = currentTime.Format("2006-01-02")
				outputSchema.DetailLogin = result
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				loginResponse.OutputSchema = outputSchema
			}
		}
		loginResponse.ErrorSchema = errorSchema
		ctx.JSON(200, loginResponse)
	}
}
