package dao

import (
	"cms_api/dbconnection"
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

type DashboardCMSDao interface {
	InquiryOverview(string, string, string) response.DashboardOverviewOutputSchema
	InquiryPromoUsage() response.DashboardPromoOutputSchema
}

type dashboardCMSDao struct {
}

func NewDashboardCMS() DashboardCMSDao {
	return &dashboardCMSDao{}
}

func returnErrOverview(code int) response.DashboardOverviewOutputSchema {
	return response.DashboardOverviewOutputSchema{User: code}
}

func returnErrPromoDashboard() response.DashboardPromoOutputSchema {
	return response.DashboardPromoOutputSchema{Objectives: []response.Objective{response.Objective{KodePromo: "-1"}}}
}

func (dao *dashboardCMSDao) InquiryOverview(chartType string, startDate string, endDate string) response.DashboardOverviewOutputSchema {
	var outputSchema response.DashboardOverviewOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	var result string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := `BEGIN GET_OVERVIEW_DATA(to_date('%s','dd-mm-yyyy'), to_date('%s','dd-mm-yyyy'),:1, :2, :3, :4, :5 ); END;`
	query = fmt.Sprintf(query, startDate, endDate)
	if _, err := db.ExecContext(ctx, query, sql.Out{Dest: &outputSchema.User}, sql.Out{Dest: &outputSchema.NewUser}, sql.Out{Dest: &outputSchema.JumlahInvestasi}, sql.Out{Dest: &outputSchema.NewPlanner}, sql.Out{Dest: &result}); err != nil {
		log.Printf("Error running %q: %+v", query, err)
		return returnErrOverview(-1)
	}
	if result == "GAGAL" {
		return returnErrOverview(-2)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrOverview(-1)
	}
	dat, err := ioutil.ReadFile(dir + "/query/OverviewChartPembelian.query")
	if err != nil {
		fmt.Println(err)
		return returnErrOverview(-1)
	}
	query = string(dat)
	query = fmt.Sprintf(query, chartType, startDate, endDate, chartType)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrOverview(-1)
	}
	defer rows.Close()
	var listBeli []response.ChartData
	for rows.Next() {
		single := response.ChartData{}
		var temp string
		rows.Scan(&temp, &single.Date, &single.Value)
		listBeli = append(listBeli, single)
	}
	outputSchema.ChartPembelian = listBeli

	dat, err = ioutil.ReadFile(dir + "/query/OverviewChartPenjualan.query")
	if err != nil {
		fmt.Println(err)
		return returnErrOverview(-1)
	}
	query = string(dat)
	query = fmt.Sprintf(query, chartType, startDate, endDate, chartType)
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrOverview(-1)
	}
	defer rows.Close()
	var listJual []response.ChartData
	for rows.Next() {
		var temp string
		single := response.ChartData{}
		rows.Scan(&temp, &single.Date, &single.Value)
		listJual = append(listJual, single)
	}
	outputSchema.ChartPenjualan = listJual
	return outputSchema
}

func (dao *dashboardCMSDao) InquiryPromoUsage() response.DashboardPromoOutputSchema {
	var outputSchema response.DashboardPromoOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrPromoDashboard()
	}
	dat, err := ioutil.ReadFile(dir + "/query/InquiryUsagePromotions.query")
	if err != nil {
		fmt.Println(err)
		return returnErrPromoDashboard()
	}
	query := string(dat)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrPromoDashboard()
	}
	defer rows.Close()
	var listPromotions []response.Promotion
	for rows.Next() {
		single := response.Promotion{}
		rows.Scan(&single.KodePromo, &single.Title, &single.Used)
		listPromotions = append(listPromotions, single)
	}
	outputSchema.Promotions = listPromotions

	dat, err = ioutil.ReadFile(dir + "/query/InquiryUsageObjectives.query")
	if err != nil {
		fmt.Println(err)
		return returnErrPromoDashboard()
	}
	query = string(dat)
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrPromoDashboard()
	}
	defer rows.Close()
	var listObjectives []response.Objective
	for rows.Next() {
		single := response.Objective{}
		rows.Scan(&single.KodePromo, &single.Title, &single.Claimed, &single.Started)
		single.Started = single.Started - single.Claimed
		listObjectives = append(listObjectives, single)
	}
	outputSchema.Objectives = listObjectives
	return outputSchema
}
