package dao

import (
	"database/sql"
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
			Id:               "-1",
			Nama:             fmt.Sprint(err),
			CurrBalance:      "-1",
			CurrProfit:       "-1",
			ChartDataOneYear: nil,
		}
		listProduk = append(listProduk, single)
		return listProduk
	}
	dat, err := ioutil.ReadFile(dir + "/query/getDailyProfitByIds.query")
	if err != nil {
		fmt.Println(err)
		single := response.GetDailyProfitWidgetOutputSchema{
			Id:               "-1",
			Nama:             fmt.Sprint(err),
			CurrBalance:      "-1",
			CurrProfit:       "-1",
			ChartDataOneYear: nil,
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
			Id:               "-1",
			Nama:             fmt.Sprint(err),
			CurrBalance:      "-1",
			CurrProfit:       "-1",
			ChartDataOneYear: nil,
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
		var chartOneYear []response.ChartData

		rows.Scan(&id, &nama, &currBalance, &currProfit)
		chartOneYear = getChartQuery(db, dir, "/query/getOneYearPerformance.query", id)

		single := response.GetDailyProfitWidgetOutputSchema{
			Id:               id,
			Nama:             nama,
			CurrBalance:      currBalance,
			CurrProfit:       currProfit,
			ChartDataOneYear: chartOneYear,
		}
		listProduk = append(listProduk, single)
	}

	return listProduk
}

func getChartQuery(db *sql.DB, dir string, filename string, id string) []response.ChartData {
	var cartData []response.ChartData
	dat, err := ioutil.ReadFile(dir + filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	query := string(dat)
	fmt.Println(query)
	query = fmt.Sprintf(query, id)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var temp1 string
		var temp2 string
		var temp3 string
		rows.Scan(&temp1, &temp2, &temp3)
		cartData = append(cartData, response.ChartData{
			DateDaily:       temp1,
			DateDailyString: temp2,
			NabDaily:        temp3,
		})
	}
	return cartData
}
