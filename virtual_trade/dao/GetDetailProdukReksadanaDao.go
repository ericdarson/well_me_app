package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"virtual_trade_api/dbconnection"
	"virtual_trade_api/entity/response"

	_ "github.com/godror/godror"
)

type GetDetailProdukReksadanaDao interface {
	GetDetailProdukReksadana(string, string) response.GetDetailProdukReksadanaOutputSchema
}

type getDetailProdukReksadanaDao struct {
}

func NewDetailProduk() GetDetailProdukReksadanaDao {
	return &getDetailProdukReksadanaDao{}
}
func returnErr() response.GetDetailProdukReksadanaOutputSchema {

	single := response.GetDetailProdukReksadanaOutputSchema{
		ID:               "-1",
		Nama:             "-1",
		Cagr:             "-1",
		Nab:              "-1",
		Expratio:         "-1",
		Aum:              "-1",
		ManagerInvestasi: "-1",
		Resiko:           "-1",
		Minimal:          "-1",
		BankKustodian:    "-1",
		BankPenampung:    "-1",
		SystemDate:       "-1",
		NabDaily:         nil,
	}

	return single
}

func (dao *getDetailProdukReksadanaDao) GetDetailProdukReksadana(idproduk string, timefilter string) response.GetDetailProdukReksadanaOutputSchema {
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErr()
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetDetailProdukReksadana.query")
	if err != nil {
		fmt.Println(err)
		return returnErr()
	}
	query := string(dat)
	query = fmt.Sprintf(query, idproduk, idproduk)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErr()
	}
	defer rows.Close()

	single := new(response.GetDetailProdukReksadanaOutputSchema)
	for rows.Next() {
		rows.Scan(&single.ID, &single.Nama, &single.Cagr, &single.Nab, &single.Expratio, &single.Aum, &single.ManagerInvestasi, &single.Resiko, &single.Minimal, &single.BankKustodian, &single.BankPenampung, &single.SystemDate)
	}
	if single.SystemDate != "" {
		layout := "02-01-2006"
		t, err := time.Parse(layout, single.SystemDate)
		if err != nil {
			fmt.Println(err)
			returnErr()
		}
		if timefilter == "oneyear" {
			t = t.AddDate(-1, 0, 0)
		} else if timefilter == "threemonths" {
			t = t.AddDate(0, -3, 0)
		} else if timefilter == "onemonth" {
			t = t.AddDate(0, -1, 0)
		} else if timefilter == "twoweeks" {
			t = t.AddDate(0, 0, -14)
		}
		temp := t.Format(layout)
		dat, err := ioutil.ReadFile(dir + "/query/GetDailyVirtualCart.query")
		if err != nil {
			fmt.Println(err)
			return returnErr()
		}
		query := string(dat)
		query = fmt.Sprintf(query, idproduk, temp, idproduk, temp)
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
			return returnErr()
		}
		defer rows.Close()
		var listDailyNab []response.DailyNab
		for rows.Next() {
			var t1 string
			var t2 string
			rows.Scan(&t1, &t2)
			dailySingle := response.DailyNab{
				Date: t1,
				Nab:  t2,
			}
			listDailyNab = append(listDailyNab, dailySingle)
		}
		single.NabDaily = listDailyNab
	}
	return *single
}
