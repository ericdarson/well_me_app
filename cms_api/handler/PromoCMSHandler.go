package handler

import (
	"cms_api/dao"
	"cms_api/entity/request"
	"cms_api/entity/response"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	daoPromo dao.PromoCMSDao = dao.NewPromoCMS()
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

func InquiryPromo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.InquiryPromoResponse
		var errorSchema response.ErrorSchema
		filter := ctx.Query("filter")
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if filter != "" && filter != "Objectives" && filter != "Promotions" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID FILTER"
			errorSchema.ErrorMessage.Indonesian = "FILTER TIDAK TERDAFTAR"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			if filter == "" {
				filter = "/query/InquiryPromo.query"
			} else if filter == "Objectives" {
				filter = "/query/InquiryPromoAkumulasi.query"
			} else if filter == "Promotions" {
				filter = "/query/InquiryPromoTransaksi.query"
			}
			outputSchema := daoPromo.InquiryPromo(filter)
			if outputSchema == nil {
				errorSchema.ErrorCode = "BIT-17-002"
				errorSchema.ErrorMessage.English = "DATA NOT FOUND"
				errorSchema.ErrorMessage.Indonesian = "DATA TIDAK DITEMUKAN"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if outputSchema[0].KodePromo == "-1" {
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
func InsertPromoAkumulasi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		layout := "02-01-2006"
		var JSONReponse response.InsertPromoAkumulasiResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.InsertPromoAkumulasiRequest
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		err := ctx.BindJSON(&JSONRequest)
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "PARAMETER REQUEST BODY TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !(re.MatchString(strings.ReplaceAll(JSONRequest.StartDate, "-", "/")) && re.MatchString(strings.ReplaceAll(JSONRequest.EndDate, "-", "/"))) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			start, err1 := time.Parse(layout, JSONRequest.StartDate)
			end, err2 := time.Parse(layout, JSONRequest.EndDate)
			if err1 != nil || err2 != nil {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
				errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if start.After(end) {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "THE STARTING DATE MUST BE EARLIER THAN THE COMPLETE DATE"
				errorSchema.ErrorMessage.Indonesian = "TANGGAL MULAI HARUS LEBIH AWAL DARI TANGGAL SELESAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else {
				outputSchema := daoPromo.InsertPromoAkumulasi(JSONRequest)
				if outputSchema.KodePromo == "-1" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(500, JSONReponse)
				} else if outputSchema.KodePromo == "-2" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
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
}

func InsertPromoTransaksi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		layout := "02-01-2006"
		var JSONReponse response.OnlyErrorSchemaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.InsertPromoTransaksiRequest
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		err := ctx.BindJSON(&JSONRequest)
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "PARAMETER REQUEST BODY TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !(re.MatchString(strings.ReplaceAll(JSONRequest.StartDate, "-", "/")) && re.MatchString(strings.ReplaceAll(JSONRequest.EndDate, "-", "/"))) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !validate_kode_promo(JSONRequest.KodePromo) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID PROMO COTE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT KODE PROMO TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !validate_kode_promo(JSONRequest.KodePromo) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID PROMO COTE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT KODE PROMO TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			start, err1 := time.Parse(layout, JSONRequest.StartDate)
			end, err2 := time.Parse(layout, JSONRequest.EndDate)
			if err1 != nil || err2 != nil {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
				errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if start.After(end) {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "THE STARTING DATE MUST BE EARLIER THAN THE COMPLETE DATE"
				errorSchema.ErrorMessage.Indonesian = "TANGGAL MULAI HARUS LEBIH AWAL DARI TANGGAL SELESAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else {
				outputSchema := daoPromo.InsertPromoTransaksi(JSONRequest)
				if outputSchema == "ERR" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(500, JSONReponse)
				} else if outputSchema == "GAGAL" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(500, JSONReponse)
				} else if outputSchema == "DUPLICATE" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "DUPLICATE VALUE FOR UNIQUE FIELD"
					errorSchema.ErrorMessage.Indonesian = "VALUE KEMBAR UNTUK FIELD YANG UNIQUE"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(401, JSONReponse)
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
}

func NonAktifPromo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var JSONReponse response.OnlyErrorSchemaResponse
		var errorSchema response.ErrorSchema
		KodePromo := ctx.Param("kode-promo")
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if KodePromo == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID INPUT PARAMETER"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			result := daoPromo.NonAktifPromo(KodePromo)
			if result == "ERROR" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if result == "GAGAL" {
				errorSchema.ErrorCode = "BIT-17-005"
				errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
				errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(500, JSONReponse)
			} else if result == "NOT FOUND" {
				errorSchema.ErrorCode = "BIT-17-004"
				errorSchema.ErrorMessage.English = "PROMO NOT FOUND OR INACTIVE"
				errorSchema.ErrorMessage.Indonesian = "PROMO TIDAK DITEMUKAN ATAU SUDAH TIDAK AKTIF"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(404, JSONReponse)
			} else if result == "SUKSES" {
				errorSchema.ErrorCode = "BIT-00-000"
				errorSchema.ErrorMessage.English = "SUCCESS"
				errorSchema.ErrorMessage.Indonesian = "BERHASIL"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(200, JSONReponse)
			}
		}
	}
}

func AdminAuthentication(ctx *gin.Context) bool {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return false
	}
	dat, err := ioutil.ReadFile(dir + "/properties/admin.properties")
	if err != nil {
		fmt.Println(err)
		return false
	}
	identity := ctx.Request.Header["Identity"]
	if identity == nil {
		return false
	} else if string(dat) == identity[0] {
		return true
	} else {
		return false
	}
}

func UpdatePromoAkumulasi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		layout := "02-01-2006"
		var JSONReponse response.OnlyErrorSchemaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.InsertPromoAkumulasiRequest
		kodePromo := ctx.Param("kode-promo")
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		err := ctx.BindJSON(&JSONRequest)
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }

		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if kodePromo == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID PARAMETER"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "PARAMETER REQUEST BODY TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !(re.MatchString(strings.ReplaceAll(JSONRequest.StartDate, "-", "/")) && re.MatchString(strings.ReplaceAll(JSONRequest.EndDate, "-", "/"))) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if strings.IndexFunc(kodePromo, isNotDigit) != -1 {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID PROMO CODE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT KODE PROMO TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			start, err1 := time.Parse(layout, JSONRequest.StartDate)
			end, err2 := time.Parse(layout, JSONRequest.EndDate)
			if err1 != nil || err2 != nil {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
				errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if start.After(end) {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "THE STARTING DATE MUST BE EARLIER THAN THE COMPLETE DATE"
				errorSchema.ErrorMessage.Indonesian = "TANGGAL MULAI HARUS LEBIH AWAL DARI TANGGAL SELESAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else {
				outputSchema := daoPromo.UpdatePromoAkumulasi(JSONRequest, kodePromo)
				if outputSchema == "ERROR" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(500, JSONReponse)
				} else if outputSchema == "GAGAL" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(500, JSONReponse)
				} else if outputSchema == "AKTIF PROMO" {
					errorSchema.ErrorCode = "BIT-17-007"
					errorSchema.ErrorMessage.English = "ACTIVE PROMO CAN'T BE UPDATED"
					errorSchema.ErrorMessage.Indonesian = "PROMO YANG SEDANG AKTIF TIDAK DAPAT DIUBAH"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(400, JSONReponse)
				} else if outputSchema == "PROMO ENDED" {
					errorSchema.ErrorCode = "BIT-17-007"
					errorSchema.ErrorMessage.English = "EXPITED PROMO CAN'T BE UPDATED"
					errorSchema.ErrorMessage.Indonesian = "PROMO YANG SUDAH TIDAK AKTIF TIDAK DAPAT DIUBAH"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(400, JSONReponse)
				} else if outputSchema == "NOT FOUND" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "PROMO NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "KODE PROMO TIDAK DITEMUKAN"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(404, JSONReponse)
				} else {
					errorSchema.ErrorCode = "BIT-00-000"
					errorSchema.ErrorMessage.English = "SUCCESS"
					errorSchema.ErrorMessage.Indonesian = "BERHASIL"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(200, JSONReponse)
				}
			}
		}
	}
}

func UpdatePromoTransaksi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		layout := "02-01-2006"
		var JSONReponse response.OnlyErrorSchemaResponse
		var errorSchema response.ErrorSchema
		var JSONRequest request.InsertPromoTransaksiRequest
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		err := ctx.BindJSON(&JSONRequest)
		kodePromo := ctx.Param("kode-promo")
		if !AdminAuthentication(ctx) {
			errorSchema.ErrorCode = "BIT-17-003"
			errorSchema.ErrorMessage.English = "ACCESS DENIED"
			errorSchema.ErrorMessage.Indonesian = "AKSES DITOLAK"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(403, JSONReponse)
		} else if kodePromo == "" {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = "INVALID PARAMETER"
			errorSchema.ErrorMessage.Indonesian = "PARAMETER TIDAK VALID"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if err != nil {
			errorSchema.ErrorCode = "BIT-17-002"
			errorSchema.ErrorMessage.English = fmt.Sprint(err)
			errorSchema.ErrorMessage.Indonesian = "PARAMETER REQUEST BODY TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !(re.MatchString(strings.ReplaceAll(JSONRequest.StartDate, "-", "/")) && re.MatchString(strings.ReplaceAll(JSONRequest.EndDate, "-", "/"))) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else if !validate_kode_promo(kodePromo) {
			errorSchema.ErrorCode = "BIT-17-006"
			errorSchema.ErrorMessage.English = "INVALID PROMO COTE FORMAT"
			errorSchema.ErrorMessage.Indonesian = "FORMAT KODE PROMO TIDAK SESUAI"
			JSONReponse.ErrorSchema = errorSchema
			ctx.JSON(400, JSONReponse)
		} else {
			start, err1 := time.Parse(layout, JSONRequest.StartDate)
			end, err2 := time.Parse(layout, JSONRequest.EndDate)
			if err1 != nil || err2 != nil {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "INVALID DATE FORMAT"
				errorSchema.ErrorMessage.Indonesian = "FORMAT TANGGAL TIDAK SESUAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else if start.After(end) {
				errorSchema.ErrorCode = "BIT-17-006"
				errorSchema.ErrorMessage.English = "THE STARTING DATE MUST BE EARLIER THAN THE COMPLETE DATE"
				errorSchema.ErrorMessage.Indonesian = "TANGGAL MULAI HARUS LEBIH AWAL DARI TANGGAL SELESAI"
				JSONReponse.ErrorSchema = errorSchema
				ctx.JSON(400, JSONReponse)
			} else {
				outputSchema := daoPromo.UpdatePromoTransaksi(JSONRequest, kodePromo)
				fmt.Println(outputSchema)

				if outputSchema == "ERROR" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(500, JSONReponse)
				} else if outputSchema == "GAGAL" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "GENERAL STORAGE ERROR"
					errorSchema.ErrorMessage.Indonesian = "SISTEM SEDANG DIPERBAIKI"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(500, JSONReponse)
				} else if outputSchema == "AKTIF PROMO" {
					errorSchema.ErrorCode = "BIT-17-007"
					errorSchema.ErrorMessage.English = "ACTIVE PROMO CAN'T BE UPDATED"
					errorSchema.ErrorMessage.Indonesian = "PROMO YANG SEDANG AKTIF TIDAK DAPAT DIUBAH"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(400, JSONReponse)
				} else if outputSchema == "PROMO ENDED" {
					errorSchema.ErrorCode = "BIT-17-007"
					errorSchema.ErrorMessage.English = "EXPITED PROMO CAN'T BE UPDATED"
					errorSchema.ErrorMessage.Indonesian = "PROMO YANG SUDAH TIDAK AKTIF TIDAK DAPAT DIUBAH"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(400, JSONReponse)
				} else if outputSchema == "NOT FOUND" {
					errorSchema.ErrorCode = "BIT-17-005"
					errorSchema.ErrorMessage.English = "PROMO NOT FOUND"
					errorSchema.ErrorMessage.Indonesian = "KODE PROMO TIDAK DITEMUKAN"
					JSONReponse.ErrorSchema = errorSchema
					ctx.JSON(404, JSONReponse)
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
}
