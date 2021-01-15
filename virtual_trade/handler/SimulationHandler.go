package handler

import (
	"regexp"
	"strings"
	"virtual_trade_api/dao"
	"virtual_trade_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoSimulation dao.SimulationDao = dao.NewSimulation()
)

func SimulationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var simulationResponse response.SimulationResponse
		var errorSchema response.ErrorSchema
		var outputSchema response.SimulationOutputSchema

		idproduk := ctx.Param("id-produk")
		simulationdate := ctx.Param("simulation-date")

		if idproduk == "" || simulationdate == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			simulationResponse.ErrorSchema = errorSchema
			ctx.JSON(400, simulationResponse)
		} else {
			re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
			if !(re.MatchString(strings.ReplaceAll(simulationdate, "-", "/"))) || len(simulationdate) != 10 {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
				errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
				simulationResponse.ErrorSchema = errorSchema
				ctx.JSON(400, simulationResponse)
			} else {
				outputSchema = daoSimulation.GetSimulationResult(idproduk, simulationdate)
				if outputSchema.NabSimulation == "-1" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					simulationResponse.ErrorSchema = errorSchema
					ctx.JSON(500, simulationResponse)
				} else if outputSchema.NabSimulation == "" {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
					simulationResponse.ErrorSchema = errorSchema
					ctx.JSON(404, simulationResponse)
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					simulationResponse.OutputSchema = outputSchema
					simulationResponse.ErrorSchema = errorSchema
					ctx.JSON(200, simulationResponse)
				}
			}
		}
	}
}
