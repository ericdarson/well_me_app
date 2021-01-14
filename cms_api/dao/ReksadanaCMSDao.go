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

type ReksadanaDao interface {
	InquiryAllJenisReksadana() []response.JenisReksadanaOutputSchema
	InsertJenisReksadana(request.InsertJenisReksadanaRequest) response.JenisReksadanaOutputSchema
	UpdateJenisReksadana(request.InsertJenisReksadanaRequest, string) response.JenisReksadanaOutputSchema
	InquiryProdukReksadana(string, string) []response.ProdukReksadanaOutputSchema
	InsertProdukReksadana(request.ProdukReksadanaRequest) string
	UpdateProdukReksadana(request.ProdukReksadanaRequest, string) string
}

type reksadanaDao struct {
}

func NewReksadana() ReksadanaDao {
	return &reksadanaDao{}
}
func returnErrInquiryJenisReksadana() []response.JenisReksadanaOutputSchema {
	return []response.JenisReksadanaOutputSchema{
		response.JenisReksadanaOutputSchema{
			ID:   "-1",
			Nama: "-1",
		},
	}
}
func (dao *reksadanaDao) InquiryAllJenisReksadana() []response.JenisReksadanaOutputSchema {
	var outputSchema []response.JenisReksadanaOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrInquiryJenisReksadana()
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetJenisReksadana.query")
	if err != nil {
		fmt.Println(err)
		return returnErrInquiryJenisReksadana()
	}
	query := string(dat)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrInquiryJenisReksadana()
	}
	defer rows.Close()

	for rows.Next() {

		var reksadana string
		var id string
		rows.Scan(&id, &reksadana)
		single := response.JenisReksadanaOutputSchema{
			ID:   id,
			Nama: reksadana,
		}
		outputSchema = append(outputSchema, single)
	}
	return outputSchema
}

func (dao *reksadanaDao) InquiryProdukReksadana(idJenis string, filter string) []response.ProdukReksadanaOutputSchema {
	var outputSchema []response.ProdukReksadanaOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return []response.ProdukReksadanaOutputSchema{response.ProdukReksadanaOutputSchema{ID: "-1"}}
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetProdukReksadana.query")
	if err != nil {
		fmt.Println(err)
		return []response.ProdukReksadanaOutputSchema{response.ProdukReksadanaOutputSchema{ID: "-1"}}
	}
	query := string(dat)
	query = fmt.Sprintf(query, filter, idJenis)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return []response.ProdukReksadanaOutputSchema{response.ProdukReksadanaOutputSchema{ID: "-1"}}
	}
	defer rows.Close()

	for rows.Next() {

		single := response.ProdukReksadanaOutputSchema{}
		rows.Scan(&single.ID, &single.Nama, &single.Nab, &single.Minimum, &single.ExpenseRatio, &single.TotalAUM, &single.ManagerInvestasi, &single.Resiko, &single.LevelResiko, &single.BankKustodian, &single.BankPenampung, &single.KinerjaSatuMinggu, &single.KinerjaSatuBulan, &single.KinerjaTigaBulan, &single.KinerjaSatuTahun, &single.IDJenis, &single.NamaJenis, &single.URLVendor, &single.PwVendor)
		outputSchema = append(outputSchema, single)
	}
	return outputSchema
}

func (dao *reksadanaDao) InsertJenisReksadana(req request.InsertJenisReksadanaRequest) response.JenisReksadanaOutputSchema {
	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_INSERT_JENIS_REKSADANA('%s','SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, req.Nama)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return response.JenisReksadanaOutputSchema{ID: "-1"}
	}
	if result != "SOMETHING WENT WRONG" {
		return response.JenisReksadanaOutputSchema{ID: result, Nama: req.Nama}
	}
	return response.JenisReksadanaOutputSchema{ID: "-2", Nama: "-2"}
}

func (dao *reksadanaDao) UpdateJenisReksadana(req request.InsertJenisReksadanaRequest, idJenis string) response.JenisReksadanaOutputSchema {
	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_UPDATE_JENIS_REKSADANA('%s','%s','SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, idJenis, req.Nama)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return response.JenisReksadanaOutputSchema{ID: "-1"}
	}
	if result != "SOMETHING WENT WRONG" {
		return response.JenisReksadanaOutputSchema{ID: idJenis, Nama: req.Nama}
	}
	return response.JenisReksadanaOutputSchema{ID: "-2", Nama: "-2"}
}

func (dao *reksadanaDao) InsertProdukReksadana(req request.ProdukReksadanaRequest) string {
	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_INSERT_PRODUK_REKSADANA('%s',%d,%f,%f,%f,'%s','%s',%d,'%s','%s','%s','%s', %f,'SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, req.Nama, req.IDJenisReksadana, req.MinimumPembelian, req.ExpenseRatio, req.TotalAUM, req.ManagerInvestasi, req.TingkatResiko, req.LevelResiko, req.BankKustodian, req.BankPenampung, req.URLVendor, req.PwVendor, req.BiayaPembelian)
	fmt.Println(query)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "ERROR"
	}
	return result
}

func (dao *reksadanaDao) UpdateProdukReksadana(req request.ProdukReksadanaRequest, idProduk string) string {
	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN SP_UPDATE_PRODUK_REKSADANA('%s', '%s',%d,%f,%f,%f,'%s','%s',%d,'%s','%s','%s','%s', %f,'SYSTEM API',:1 ); END;`
	query = fmt.Sprintf(query, idProduk, req.Nama, req.IDJenisReksadana, req.MinimumPembelian, req.ExpenseRatio, req.TotalAUM, req.ManagerInvestasi, req.TingkatResiko, req.LevelResiko, req.BankKustodian, req.BankPenampung, req.URLVendor, req.PwVendor, req.BiayaPembelian)
	fmt.Println(query)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return "ERROR"
	}
	return result
}
