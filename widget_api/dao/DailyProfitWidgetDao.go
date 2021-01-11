package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"widget_api/dbconnection"
	"widget_api/entity/response"

	_ "github.com/godror/godror"
)

type DailyProfitWidgetDao interface {
	GetDailyProfitWidget(string) []response.GetDailyProfitWidgetOutputSchema
}

type dailyProfitWidgetDao struct {
	temp string
}

func NewDailyProfit() DailyProfitWidgetDao {
	return &dailyProfitWidgetDao{}
}

func (dao *dailyProfitWidgetDao) GetDailyProfitWidget(listIds string) []response.GetDailyProfitWidgetOutputSchema {
	var listProduk []response.GetDailyProfitWidgetOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		single := response.GetDailyProfitWidgetOutputSchema{
			Id:          "-1",
			Nama:        fmt.Sprint(err),
			CurrBalance: "-1",
			CurrProfit:  "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	dat, err := ioutil.ReadFile(dir + "/query/getDailyProfitByIds.query")
	if err != nil {
		fmt.Println(err)
		single := response.GetDailyProfitWidgetOutputSchema{
			Id:          "-1",
			Nama:        fmt.Sprint(err),
			CurrBalance: "-1",
			CurrProfit:  "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	query := string(dat)
	query = fmt.Sprintf(query, listIds)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		single := response.GetDailyProfitWidgetOutputSchema{
			Id:          "-1",
			Nama:        fmt.Sprint(err),
			CurrBalance: "-1",
			CurrProfit:  "-1",
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var nama string
		var currBalance string
		var currProfit string
		rows.Scan(&id, &nama, &currBalance, &currProfit)
		single := response.GetDailyProfitWidgetOutputSchema{
			Id:          id,
			Nama:        nama,
			CurrBalance: currBalance,
			CurrProfit:  currProfit,
		}
		listProduk = append(listProduk, single)
	}

	return listProduk
}
