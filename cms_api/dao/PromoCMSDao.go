package dao

import (
	"cms_api/dbconnection"
	"cms_api/entity/request"
	"cms_api/entity/response"
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/godror/godror"
)

type PromoCMSDao interface {
	InsertPromoAkumulasi(request.InsertPromoAkumulasiRequest) response.InsertPromoAkumulasiOutputSchema
	InsertPromoTransaksi(request.InsertPromoTransaksiRequest) string
	NonAktifPromo(string) string
	InquiryPromo(string) []response.InquiryPromoOutputSchema
	UpdatePromoAkumulasi(request.InsertPromoAkumulasiRequest, string) string
	UpdatePromoTransaksi(request.InsertPromoTransaksiRequest, string) string
}

type promoCMSDao struct {
}

func NewPromoCMS() PromoCMSDao {
	return &promoCMSDao{}
}

func returnErrInsertPromoAkumulasi() response.InsertPromoAkumulasiOutputSchema {
	return response.InsertPromoAkumulasiOutputSchema{
		KodePromo:   "-1",
		Title:       "-1",
		Subtitle:    "-1",
		StartDate:   "-1",
		EndDate:     "-1",
		Description: "-1",
		Target:      -1,
		CashBack:    -1,
	}
}

func returnErrInquiryPromo() []response.InquiryPromoOutputSchema {
	return []response.InquiryPromoOutputSchema{response.InquiryPromoOutputSchema{
		KodePromo:    "-1",
		Title:        "-1",
		Subtitle:     "-1",
		Description:  "-1",
		Target:       "-1",
		Cashback:     "-1",
		Minimum:      "-1",
		StartDate:    "-1",
		EndDate:      "-1",
		PromoType:    "-1",
		ActiveStatus: "-1",
	}}
}

func (dao *promoCMSDao) InquiryPromo(filename string) []response.InquiryPromoOutputSchema {

	var outputSchema []response.InquiryPromoOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrInquiryPromo()
	}
	dat, err := ioutil.ReadFile(dir + filename)
	if err != nil {
		fmt.Println(err)
		return returnErrInquiryPromo()
	}
	query := string(dat)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrInquiryPromo()
	}
	defer rows.Close()

	for rows.Next() {
		single := response.InquiryPromoOutputSchema{}
		rows.Scan(&single.KodePromo, &single.Title, &single.Subtitle, &single.Description, &single.Target, &single.Cashback, &single.Minimum, &single.StartDate, &single.EndDate, &single.PromoType, &single.ActiveStatus)
		outputSchema = append(outputSchema, single)
	}

	return outputSchema
}

func (dao *promoCMSDao) InsertPromoAkumulasi(req request.InsertPromoAkumulasiRequest) response.InsertPromoAkumulasiOutputSchema {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var outputSchema response.InsertPromoAkumulasiOutputSchema
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_ADD_PROMO_AKUMULASI('%s','%s', to_date('%s','dd-mm-yyyy'), to_date('%s','dd-mm-yyyy'),'%s', %f, %f ,'SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, req.Title, req.Subtitle, req.StartDate, req.EndDate, req.Description, req.CashBack, req.Target)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return returnErrInsertPromoAkumulasi()
	}

	if result == "GAGAL" {
		ret := returnErrInsertPromoAkumulasi()
		ret.KodePromo = "-2"
	} else {
		outputSchema.KodePromo = result
		outputSchema.Title = req.Title
		outputSchema.Subtitle = req.Subtitle
		outputSchema.StartDate = req.StartDate
		outputSchema.EndDate = req.EndDate
		outputSchema.Description = req.Description
		outputSchema.CashBack = req.CashBack
		outputSchema.Target = req.Target
	}
	return outputSchema
}

func (dao *promoCMSDao) InsertPromoTransaksi(req request.InsertPromoTransaksiRequest) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_ADD_PROMO_WITH_KODE('%s','%s','%s', to_date('%s','dd-mm-yyyy'), to_date('%s','dd-mm-yyyy'),'%s', %f, %f ,'SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, req.KodePromo, req.Title, req.Subtitle, req.StartDate, req.EndDate, req.Description, req.CashBack, req.Minimum)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "ERROR"
	}

	return result
}

func (dap *promoCMSDao) NonAktifPromo(kode_promo string) string {
	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_NONAKTIF_PROMO('%s','SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, kode_promo)
	fmt.Println(query)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "ERROR"
	}

	return result
}

func (dao *promoCMSDao) UpdatePromoAkumulasi(req request.InsertPromoAkumulasiRequest, kodePromo string) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_UPDATE_PROMO_AKUMULASI('%s', '%s','%s', to_date('%s','dd-mm-yyyy'), to_date('%s','dd-mm-yyyy'),'%s', %f, %f ,'SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, kodePromo, req.Title, req.Subtitle, req.StartDate, req.EndDate, req.Description, req.CashBack, req.Target)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "ERROR"
	}

	return result
}

func (dao *promoCMSDao) UpdatePromoTransaksi(req request.InsertPromoTransaksiRequest, kodePromo string) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_UPDATE_PROMO_WITH_KODE('%s','%s','%s', to_date('%s','dd-mm-yyyy'), to_date('%s','dd-mm-yyyy'),'%s', %f, %f ,'SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, kodePromo, req.Title, req.Subtitle, req.StartDate, req.EndDate, req.Description, req.CashBack, req.Minimum)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "ERROR"
	}

	return result
}
