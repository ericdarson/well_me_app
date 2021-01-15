package handler

import (
	"pembelian_verifikasi_bank_api/dao"
	"pembelian_verifikasi_bank_api/entity/response"

	"github.com/gin-gonic/gin"
)

var (
	daoPlannerVerifikasiBankPembelian dao.PlannerVerifikasiBankPembelianDao = dao.NewPlannerVerifikasiBankPembelian()
)

func PlannerVerifikasiBankPembelianHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.PlannerVerifikasiBankPembelianResponse
		var errorSchema response.ErrorSchema

		vendorPw := ctx.Request.Header["Identity"]
		idTrans := ctx.Param("id-transaksi")
		if vendorPw == nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "MISSING REQUIRED HEADER FIELD"
			errorSchema.ErrorMessage.Indonesian = "HEADER YANG DIPERLUKAN BELUM DIISI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if idTrans == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID TRANSACTION ID FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT ID TRANSAKSI TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoPlannerVerifikasiBankPembelian.DoVerifikasi(vendorPw[0], idTrans)
			if outputSchema == "WRONG PASSWORD" {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "FORBIDDEN ACCESS"
				errorSchema.ErrorMessage.Indonesian = "TIDAK MEMILIKI HAK AKSES"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(403, JSONReponse)
			} else if outputSchema == "SYSTEM ERROR" || outputSchema == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "NOT FOUND" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "TRANSACTION NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "TRANSAKSI TIDAK DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
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
