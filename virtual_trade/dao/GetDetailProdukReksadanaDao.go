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
		Expratio:         "-1",
		Aum:              "-1",
		ManagerInvestasi: "-1",
		Resiko:           "-1",
		Minimal:          "-1",
		BankKustodian:    "-1",
		BankPenampung:    "-1",
		SystemDate:       "-1",
		SystemDateString: "-1",
	}

	return single
}

func (dao *getDetailProdukReksadanaDao) GetDetailProdukReksadana(idproduk string, simulationDate string) response.GetDetailProdukReksadanaOutputSchema {
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
	query = fmt.Sprintf(query, simulationDate, idproduk)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErr()
	}
	defer rows.Close()

	single := response.GetDetailProdukReksadanaOutputSchema{}
	for rows.Next() {
		rows.Scan(&single.ID, &single.Nama, &single.Nab, &single.Expratio, &single.Aum, &single.ManagerInvestasi, &single.Resiko, &single.Minimal, &single.BankKustodian, &single.BankPenampung, &single.SystemDate, &single.SystemDateString)
	}
	if single.SystemDate != "" {
		layout := "02-01-2006"
		t, err := time.Parse(layout, single.SystemDate)
		if err != nil {
			fmt.Println(err)
			returnErr()
		}
		var arrTime []time.Time
		var arrFilename []string
		arrTime = append(arrTime, t.AddDate(0, 0, -7))
		arrTime = append(arrTime, t.AddDate(0, -1, 0))
		arrTime = append(arrTime, t.AddDate(0, -3, 0))
		arrTime = append(arrTime, t.AddDate(-1, 0, 0))

		arrFilename = append(arrFilename, "/query/GetDailyVirtualCart.query")
		arrFilename = append(arrFilename, "/query/GetDailyVirtualCart.query")
		arrFilename = append(arrFilename, "/query/GetDailyVirtualCart.query")
		arrFilename = append(arrFilename, "/query/GetDailyVirtualCartOneYear.query")

		var arrListNab [][]response.DailyNab
		for i := 0; i < 4; i++ {
			temp := arrTime[i].Format(layout)
			dat, err := ioutil.ReadFile(dir + arrFilename[i])
			if err != nil {
				fmt.Println(err)
				return returnErr()
			}
			query := string(dat)
			query = fmt.Sprintf(query, idproduk, temp)
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
				var t3 float64
				rows.Scan(&t1, &t2, &t3)
				dailySingle := response.DailyNab{
					Date:       t1,
					DateString: t2,
					Nab:        t3,
				}
				listDailyNab = append(listDailyNab, dailySingle)
			}
			arrListNab = append(arrListNab, listDailyNab)
		}
		single.NabOneWeek = arrListNab[0]
		single.NabOneMonth = arrListNab[1]
		single.NabThreeMonths = arrListNab[2]
		single.NabOneYear = arrListNab[3]

		single.CagrOneWeek = (single.Nab - single.NabOneWeek[0].Nab) / single.NabOneWeek[0].Nab * 100
		single.CagrOneMonth = (single.Nab - single.NabOneMonth[0].Nab) / single.NabOneMonth[0].Nab * 100
		single.CagrThreeMonths = (single.Nab - single.NabThreeMonths[0].Nab) / single.NabThreeMonths[0].Nab * 100
		single.CagrOneYear = (single.Nab - single.NabOneYear[0].Nab) / single.NabOneYear[0].Nab * 100
	}

	return single
}
