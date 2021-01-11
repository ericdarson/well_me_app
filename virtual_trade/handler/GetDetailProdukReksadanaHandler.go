package handler

import (
	"regexp"
	"strings"
	"virtual_trade_api/dao"
	"virtual_trade_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoDetailProduk dao.GetDetailProdukReksadanaDao = dao.NewDetailProduk()
)

func GetDetailProdukReksadanaHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idproduk := ctx.Param("id-produk")
		simulationdate := ctx.Param("simulation-date")
		var detailProdukReksadanaResponse response.GetDetailProdukReksadanaResponse
		var outputSchema response.GetDetailProdukReksadanaOutputSchema
		var errorSchema response.ErrorSchema

		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		b := strings.IndexFunc(idproduk, isNotDigit) == -1
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		if idproduk == "" || simulationdate == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETERS"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			detailProdukReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, detailProdukReksadanaResponse)
		} else if !b {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "ID MUST ONLY CONTAIN A NUMBER"
			errorSchema.ErrorMessage.Indonesian = "ID HANYA BOLEH MENGANDUNG ANGKA"
			detailProdukReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, detailProdukReksadanaResponse)
		} else if !(re.MatchString(strings.ReplaceAll(simulationdate, "-", "/"))) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
			detailProdukReksadanaResponse.ErrorSchema = errorSchema
			ctx.JSON(400, detailProdukReksadanaResponse)
		} else {
			result := daoDetailProduk.GetDetailProdukReksadana(idproduk, simulationdate)
			if result.ID == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				detailProdukReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(500, detailProdukReksadanaResponse)
			} else if result.ID == "" {
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				detailProdukReksadanaResponse.ErrorSchema = errorSchema
				ctx.JSON(404, detailProdukReksadanaResponse)
			} else {
				outputSchema = result
				if outputSchema.ID == "" {

				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					detailProdukReksadanaResponse.OutputSchema = outputSchema
					detailProdukReksadanaResponse.ErrorSchema = errorSchema
					ctx.JSON(200, detailProdukReksadanaResponse)
				}
			}
		}
	}
}
