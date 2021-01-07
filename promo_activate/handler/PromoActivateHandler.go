package handler

import (
	"fmt"
	"promo_activate/dao"
	"promo_activate/entity/request"
	"promo_activate/entity/response"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	daoActivatePromo dao.PromoActivateDao = dao.NewPromoActivate()
)

func ActivatePromo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.ActivatePromoResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.ActivatePromoRequest
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		err := ctx.BindJSON(&JSONRequest)
		if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "PARAMETER REQUEST BODY TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !(strings.IndexFunc(JSONRequest.KodePromo, isNotDigit) == -1) || len(JSONRequest.KodePromo) != 10 {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID PROMO CODE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT KODE PROMO TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoActivatePromo.ActivatePromo(JSONRequest)
			if outputSchema == "SYSTEM ERROR" || outputSchema == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "NOT FOUND" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PROMO CODE NOT FOUND OR EXPIRED"
				errorSchema.ErrorMessage.Indonesian = "KODE PROMO TIDAK DITEMUKAN ATAU PERIODE PROMO SUDAH BERAKHIR"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if outputSchema == "ALREADY CLAIMED" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PROMO CODE ALREADY USED"
				errorSchema.ErrorMessage.Indonesian = "KODE PROMO SUDAH DIGUNAKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(403, JSONReponse)
			} else if outputSchema == "SUKSES" {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
			}
		}
	}

}

func InquiryPromo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bcaID := ctx.Param("bca-id")

		var JSONReponse response.ListPromoResposne
		var errorSchema response.ErrorSchema
		if bcaID == "" {
			errorSchema.ErrorCode = "BIT-17-005"
			errorSchema.ErrorMessage.English = "BCA ID IS REQUIRED"
			errorSchema.ErrorMessage.Indonesian = "BCA ID WAJIB DIISI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoActivatePromo.InquiryPromo(bcaID)
			if outputSchema.Objectives == nil && outputSchema.Promotion == nil {
				errorSchema.ErrorCode = "BIT-17-002"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if outputSchema.Objectives[0].KodePromo == "-1" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				JSONReponse.OutputSchema = outputSchema
				ctx.JSON(200, JSONReponse)
			}
		}
	}
}

func ClaimPromo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.ActivatePromoResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.ActivatePromoRequest
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		err := ctx.BindJSON(&JSONRequest)
		if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "PARAMETER REQUEST BODY TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !(strings.IndexFunc(JSONRequest.KodePromo, isNotDigit) == -1) || len(JSONRequest.KodePromo) != 10 {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID PROMO CODE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT KODE PROMO TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			outputSchema := daoActivatePromo.ClaimPromo(JSONRequest)
			if outputSchema == "SYSTEM ERROR" || outputSchema == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if outputSchema == "NOT FOUND" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PROMO CODE NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "KODE PROMO TIDAK DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if outputSchema == "ALREADY CLAIMED" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PROMO HAVEN'T REACH DESIRED VALUE OR ALREADY BEEN USED"
				errorSchema.ErrorMessage.Indonesian = "PROMO BELUM MENCAPAI TARGET AKUMULASI ATAU SUDAH DIGUNAKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if outputSchema == "EXPIRED" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "PROMO HAS ENDED"
				errorSchema.ErrorMessage.Indonesian = "PERIODE PROMO TELAH HABIS"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else if outputSchema == "SUKSES" {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			} else {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
			}
		}
	}

}
