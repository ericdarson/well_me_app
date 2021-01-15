package handler

import (
	"regexp"
	"strings"
	"virtual_trade_api/dao"
	"virtual_trade_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoProjection dao.ProjectionDao = dao.NewProjection()
)

func ProjectionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var projectionResponse response.ProjectionResponse
		var errorSchema response.ErrorSchema
		idproduk := ctx.Param("id-produk")
		simulationdate := ctx.Param("simulation-date")
		if idproduk == "" || simulationdate == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			projectionResponse.ErrorSchema = errorSchema
			ctx.JSON(400, projectionResponse)
		} else {
			re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
			if !(re.MatchString(strings.ReplaceAll(simulationdate, "-", "/"))) || len(simulationdate) != 10 {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
				errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
				projectionResponse.ErrorSchema = errorSchema
				ctx.JSON(400, projectionResponse)
			} else {
				outputSchema := daoProjection.GetProjectionResult(idproduk, simulationdate)
				if outputSchema.ID == "-1" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					projectionResponse.ErrorSchema = errorSchema
					ctx.JSON(500, projectionResponse)
				} else if outputSchema.ID == "-2" {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
					projectionResponse.ErrorSchema = errorSchema
					ctx.JSON(404, projectionResponse)
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					projectionResponse.ErrorSchema = errorSchema
					projectionResponse.OutputSchema = outputSchema
					ctx.JSON(200, projectionResponse)
				}
			}
		}
	}
}
