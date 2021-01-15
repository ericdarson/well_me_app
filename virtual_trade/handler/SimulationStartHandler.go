package handler

import (
	"regexp"
	"strings"
	"virtual_trade_api/dao"
	"virtual_trade_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoSimulationStart dao.SimulationStartDao = dao.NewSimulationStart()
)

func SimulationStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idproduk := ctx.Param("id-produk")
		jumlahinvest := ctx.Param("jumlahinvest")
		startDate := ctx.Param("starting-date")
		var errorSchema response.ErrorSchema
		var outputSchema response.SimulationStartOutputSchema
		var simulationStartResponse response.SimulationStartResponse
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		if idproduk == "" || jumlahinvest == "" || startDate == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			simulationStartResponse.ErrorSchema = errorSchema
			ctx.JSON(400, simulationStartResponse)
		} else if !(re.MatchString(strings.ReplaceAll(startDate, "-", "/"))) || len(startDate) != 10 {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
			simulationStartResponse.ErrorSchema = errorSchema
			ctx.JSON(400, simulationStartResponse)
		} else {
			isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
			b := strings.IndexFunc(jumlahinvest, isNotDigit) == -1
			if !b {
				errorSchema.ErrorCode = "BIT-17-002"
				errorSchema.ErrorMessage.English = "NUMBER PARAMETERS CONTAIN TEXT"
				errorSchema.ErrorMessage.Indonesian = "INPUTAN ANGKA MENGANDUNG KARAKTER"
				simulationStartResponse.ErrorSchema = errorSchema
				ctx.JSON(400, simulationStartResponse)
			} else {
				outputSchema = daoSimulationStart.StartSimulation(idproduk, jumlahinvest, startDate)
				if outputSchema.StartDate == "-1" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					simulationStartResponse.ErrorSchema = errorSchema
					ctx.JSON(500, simulationStartResponse)
				} else if outputSchema.StartDate == "-2" {
					errorSchema.ErrorCode = "BIT-10-001"
					errorSchema.ErrorMessage.English = "DO NOT REACH THE MINIMUM TRANSACTION"
					errorSchema.ErrorMessage.Indonesian = "TIDAK MENCAPAI MINIMUM TRANSAKSI"
					simulationStartResponse.ErrorSchema = errorSchema
					ctx.JSON(200, simulationStartResponse)
				} else if outputSchema.StartDate == "" {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
					simulationStartResponse.ErrorSchema = errorSchema
					ctx.JSON(404, simulationStartResponse)
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					simulationStartResponse.OutputSchema = outputSchema
					simulationStartResponse.ErrorSchema = errorSchema
					ctx.JSON(200, simulationStartResponse)
				}
			}
		}
	}
}

/*
func SimulationStartHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startingdate := ctx.Param("starting-date")
		simulationdate := ctx.Param("simulation-date")

		var simulationResponse response.SimulationResponse
		var outputSchema response.SimulationOutputSchema
		var errorSchema response.ErrorSchema

		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/\\d\\d\\d\\d)")

		layout := "02-01-2006"
		if !(re.MatchString(startingdate) && re.MatchString(simulationdate)) {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
		} else if (time.Parse(layout, startingdate)).After(time.Parse(layout, simulationdate)) {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
		} else {
			result := daoSimulation.GetSimulationData(startingdate, simulationdate)
			if result[0].Id == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
			} else {
				outputSchema = result
				if outputSchema == nil {
					errorSchema.ErrorCode = "BIT-17-004"
					errorSchema.ErrorMessage.English = "DATA NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				}
				getListProdukReksadana.OutputSchema = outputSchema
			}
		}

		simulationResponse.ErrorSchema = errorSchema
		ctx.JSON(200, simulationResponse)
	}
}
*/
