package handler

import (
	"fmt"
	"planner_pembelian_api/dao"
	"planner_pembelian_api/entity/request"
	"planner_pembelian_api/entity/response"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	daoPlannerPembelian dao.PlannerPembelianDao = dao.NewPlannerPembelian()
)

func validate_kode_promo(kode_promo string) bool {
	if kode_promo == "" {
		return true
	} else {
		isNotAlphaNumeric := func(c rune) bool { return (c < '0' || c > '9') && (c < 'A' || c > 'Z') }
		if strings.IndexFunc(kode_promo, isNotAlphaNumeric) == -1 {
			for _, char := range kode_promo {
				if char >= 'A' && char <= 'Z' {
					return true
				}
			}
		}
		return false
	}
}

func PlannerPembelianHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.PlannerPembelianResponse
		var errorSchema response.ErrorSchema

		var JSONRequest request.PlannerPembelianRequest
		err := ctx.BindJSON(&JSONRequest)

		if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "PARAMETER INPUT TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !validate_kode_promo(JSONRequest.KodePromo) {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID PROMO CODE"
			errorSchema.ErrorMessage.Indonesian = "KODE PROMO TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoPlannerPembelian.DoPembelian(JSONRequest)
			fmt.Println(outputSchema)
			if outputSchema == "PROMO NOT REACH MINIMUM" {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "HAVEN'T REACH MINIMUM PURCHASE"
				errorSchema.ErrorMessage.Indonesian = "BELUM MENCAPAI MINIMAL PEMBELIAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if outputSchema == "SYSTEM ERROR" || outputSchema == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "RISK PROFILE NOT MATCH" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PRODUCTS DO NOT MATCH THE RISK PROFILE"
				errorSchema.ErrorMessage.Indonesian = "PRODUK TIDAK SESUAI PROFIL RISIKO"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "FAILED TRANSACTION" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "INCOMPLETE TRANSACTION"
				errorSchema.ErrorMessage.Indonesian = "TRANSAKSI TIDAK BERHASIL, SILAHKAN HUBUNGI HALOBCA"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "KODE PROMO NOT FOUND" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PROMO CODE NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "KODE PROMO TIDAK ADA"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "OUT OF DATE" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "THE PROMO HAS ENDED"
				errorSchema.ErrorMessage.Indonesian = "PERIODE PROMO TELAH HABIS"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if outputSchema == "USED PROMO" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PROMO ALREADY USED"
				errorSchema.ErrorMessage.Indonesian = "PROMO SUDAH DIGUNAKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if outputSchema == "TIME OUT" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "CONNECTION TIME OUT"
				errorSchema.ErrorMessage.Indonesian = "KONEKSI BERMASALAH"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if outputSchema == "SALDO TIDAK CUKUP" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "INSUFFICIENT BALANCE"
				errorSchema.ErrorMessage.Indonesian = "SALDO REKENING ANDA TIDAK CUKUP TRANSAKSI DIBATALKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if outputSchema == "SUKSES" {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			}
		}
	}
}
