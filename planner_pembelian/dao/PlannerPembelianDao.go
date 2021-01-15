package dao

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"planner_pembelian_api/dbconnection"
	"planner_pembelian_api/entity"
	"planner_pembelian_api/entity/request"
	"strconv"
	"time"

	_ "github.com/godror/godror"
)

type PlannerPembelianDao interface {
	DoPembelian(request.PlannerPembelianRequest) string
}

type plannerPembelianDao struct {
}

func NewPlannerPembelian() PlannerPembelianDao {
	return &plannerPembelianDao{}
}

func (dao *plannerPembelianDao) DoPembelian(req request.PlannerPembelianRequest) string {

	conn := dbconnection.New()
	db := conn.GetConnection()
	var result string
	var totalTransaksi float64
	for _, single := range req.Products {
		totalTransaksi += single.Nominal
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if req.KodePromo != "" {
		query := `BEGIN SP_USE_KODEPROMO('%s','%s',%f,'SYSTEM API',1,:1 ); END;`
		query = fmt.Sprintf(query, req.BcaID, req.KodePromo, totalTransaksi)
		if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
			log.Printf("Error running %q: %+v", query, err)
			return "TIME OUT"
		}
		if result != "SUKSES" {
			return result
		}
	}
	var listBiayaTransaksi []float64
	var totalSaldoBerkurang float64
	totalSaldoBerkurang = 0
	for _, single := range req.Products {
		var biayaPembelian float64
		query := `BEGIN SP_VALIDATE_AND_GET_BIAYA('%s',%d,%f,:1 ); END;`
		query = fmt.Sprintf(query, req.BcaID, single.IDProduk, single.Nominal)
		if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &biayaPembelian}); err != nil {
			log.Printf("Error running %q: %+v", query, err)
			return "TIME OUT"
		}
		if biayaPembelian == -1 || biayaPembelian == -2 || biayaPembelian == -3 {
			query = `BEGIN SP_CANCEL_USE_PROMO('%s','%s); END;`
			query = fmt.Sprintf(query, req.BcaID, req.KodePromo)
			if _, err := db.ExecContext(ctx, query); err != nil {
				log.Printf("Error running %q: %+v", query, err)
				return "TIME OUT"
			}
			if biayaPembelian == -1 {
				return "UNDER MINIMUM"
			} else if biayaPembelian == -2 {
				return "RISK PROFILE NOT MATCH"
			} else {
				return "GAGAL"
			}
		}
		listBiayaTransaksi = append(listBiayaTransaksi, biayaPembelian)
		totalSaldoBerkurang += biayaPembelian + single.Nominal
	}
	var nomorRekening string
	query := "select NO_REKENING from T_NASABAH where BCA_ID = '" + req.BcaID + "'"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return "TIME OUT"
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nomorRekening)
	}
	if !checkSaldoToMainFrame(nomorRekening, totalSaldoBerkurang) {
		query := `BEGIN SP_CANCEL_USE_PROMO('%s','%s'); END;`
		query = fmt.Sprintf(query, req.BcaID, req.KodePromo)
		if _, err := db.ExecContext(ctx, query); err != nil {
			log.Printf("Error running %q: %+v", query, err)
			return "TIME OUT"
		}
		return "SALDO TIDAK CUKUP"
	}
	if !mainFrameTransaction(db, nomorRekening, req.Products, listBiayaTransaksi, totalSaldoBerkurang) {
		query := `BEGIN SP_CANCEL_USE_PROMO('%s','%s'); END;`
		query = fmt.Sprintf(query, req.BcaID, req.KodePromo)
		if _, err := db.ExecContext(ctx, query); err != nil {
			log.Printf("Error running %q: %+v", query, err)
			return "TIME OUT"
		}
		return "FAILED TRANSACTION"
	}
	for _, single := range req.Products {
		query := `BEGIN SP_REQUEST_PEMBELIAN('%s',%d,%d,'%f','%s', 'SYSTEM API',1,:1 ); END;`
		query = fmt.Sprintf(query, req.BcaID, single.IDProduk, req.IDPlan, single.Nominal, req.KodePromo)
		if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &result}); err != nil {
			log.Printf("Error running %q: %+v", query, err)
			return "TIME OUT"
		}
		if result != "SUKSES" {
			query = `BEGIN ROLLBACK; END;`
			if _, err := db.ExecContext(ctx, query); err != nil {
				log.Printf("Error running %q: %+v", query, err)
				return "TIME OUT"
			}
			return result
		}
	}

	return result
}

func checkSaldoToMainFrame(nomorRekening string, totalSaldoBerkurang float64) bool {
	//start HIT mainframe
	url := "http://10.20.218.9:9083/project-wm02/inqSaldo?rek-no=" + nomorRekening
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return false
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var responseJSON entity.GetSaldo
		json.Unmarshal(data, &responseJSON)
		if responseJSON.Reponse.ErrorResponse.ErrorCode == "BIT-17-000" {
			return totalSaldoBerkurang <= responseJSON.Reponse.Response.Saldo
		} else {
			return false
		}
	}
	//end HIT mainframe

}

func mainFrameTransaction(db *sql.DB, nomorRekening string, products []request.ProdukBeli, listBiayaTransaksi []float64, totalSaldoBerkurang float64) bool {
	var listNamaProduk []string
	for _, product := range products {
		var namaProduct string

		query := "select NAMA_PRODUK from T_PRODUK_REKSADANA where ID_PRODUK = '" + strconv.Itoa(product.IDProduk) + "'"
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
			return false
		}
		for rows.Next() {
			rows.Scan(&namaProduct)
		}
		listNamaProduk = append(listNamaProduk, namaProduct)
	}

	reqMainFrame := entity.TransactionMFRequest{}
	reqMainFrame.Parent.Response.User.FuncID = "1"
	reqMainFrame.Parent.Response.User.NoRek = nomorRekening
	reqMainFrame.Parent.Response.User.TotalNominal = totalSaldoBerkurang
	for i := 0; i < len(listBiayaTransaksi); i++ {
		single := entity.TransactionMFRequestListTrn{
			Desc:    "Pembelian Reksadana " + listNamaProduk[i],
			Nominal: listBiayaTransaksi[i] + products[i].Nominal,
		}
		reqMainFrame.Parent.Response.User.ListTrans = append(reqMainFrame.Parent.Response.User.ListTrans, single)
	}
	jsonValue, err := json.Marshal(reqMainFrame)
	if err != nil {
		fmt.Println(err)
		return false
	}
	request, err := http.NewRequest("POST", "http://10.20.218.9:9083/project-wm06/trans", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return false
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var responseJSON entity.TransactionMF
		json.Unmarshal(data, &responseJSON)

		return responseJSON.Reponse.ErrorResponse.ErrorCode == "BIT-17-000"
	}
}
