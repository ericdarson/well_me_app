package dao

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"
	"virtual_trade_api/dbconnection"
	"virtual_trade_api/entity/response"

	_ "github.com/godror/godror"
)

type ProjectionDao interface {
	GetProjectionResult(string, string) response.ProjectionOutputSchema
}

type projectionDao struct {
}

func NewProjection() ProjectionDao {
	return &projectionDao{}
}
func returnErrProjection() response.ProjectionOutputSchema {
	return response.ProjectionOutputSchema{
		ID:                  "-1",
		Nama:                "-1",
		Date:                "-1",
		DateString:          "-1",
		Nab:                 "-1",
		CagrOneYear:         -1,
		CagrThreeMonths:     -1,
		CagrOneMonth:        -1,
		CagrOneWeek:         -1,
		CartDataOneYear:     nil,
		CartDataThreeMonths: nil,
		CartDataOneMonth:    nil,
		CartDataOneWeek:     nil,
	}
}

func (dao *projectionDao) GetProjectionResult(idproduk string, datesimulation string) response.ProjectionOutputSchema {
	var projectionData response.ProjectionOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrProjection()
	}

	dat, err := ioutil.ReadFile(dir + "/query/ProjectionDetail.query")
	if err != nil {
		fmt.Println(err)
		return returnErrProjection()
	}
	query := string(dat)
	query = fmt.Sprintf(query, idproduk, datesimulation)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrProjection()
	}
	defer rows.Close()
	boolErr := true
	for rows.Next() {
		boolErr = false
		rows.Scan(&projectionData.ID, &projectionData.Nama, &projectionData.Date, &projectionData.DateString, &projectionData.Nab)
	}
	if boolErr {
		return response.ProjectionOutputSchema{
			ID:                  "-2",
			Nama:                "-2",
			Date:                "-2",
			DateString:          "-2",
			Nab:                 "-2",
			CagrOneYear:         -2,
			CagrThreeMonths:     -2,
			CagrOneMonth:        -2,
			CagrOneWeek:         -2,
			CartDataOneYear:     nil,
			CartDataThreeMonths: nil,
			CartDataOneMonth:    nil,
			CartDataOneWeek:     nil,
		}
	} else {
		layout := "02-01-2006"
		t, err := time.Parse(layout, datesimulation)
		if err != nil {
			returnErrProjection()
		}
		toneyear := (t.AddDate(-1, 0, 0)).Format(layout)
		tthreemonths := (t.AddDate(0, -3, 0)).Format(layout)
		tonemonth := (t.AddDate(0, -1, 0)).Format(layout)
		toneweek := (t.AddDate(0, 0, -7)).Format(layout)
		projectionData.CartDataOneWeek = getCartQuery(db, dir, "/query/ProjectionCart.query", idproduk, datesimulation, toneweek)
		projectionData.CartDataOneMonth = getCartQuery(db, dir, "/query/ProjectionCart.query", idproduk, datesimulation, tonemonth)
		projectionData.CartDataThreeMonths = getCartQuery(db, dir, "/query/ProjectionCart.query", idproduk, datesimulation, tthreemonths)
		projectionData.CartDataOneYear = getCartQuery(db, dir, "/query/ProjectionCartOneYear.query", idproduk, datesimulation, toneyear)

		nabfloat, _ := strconv.ParseFloat(projectionData.Nab, 64)
		projectionData.CagrOneWeek, _ = strconv.ParseFloat(projectionData.CartDataOneWeek[0].NabDaily, 64)
		projectionData.CagrOneWeek = math.Round(((nabfloat-projectionData.CagrOneWeek)/projectionData.CagrOneWeek*100)*100) / 100

		projectionData.CagrOneMonth, _ = strconv.ParseFloat(projectionData.CartDataOneMonth[0].NabDaily, 64)
		projectionData.CagrOneMonth = math.Round(((nabfloat-projectionData.CagrOneMonth)/projectionData.CagrOneMonth*100)*100) / 100

		projectionData.CagrThreeMonths, _ = strconv.ParseFloat(projectionData.CartDataThreeMonths[0].NabDaily, 64)
		projectionData.CagrThreeMonths = math.Round(((nabfloat-projectionData.CagrThreeMonths)/projectionData.CagrThreeMonths*100)*100) / 100

		projectionData.CagrOneYear, _ = strconv.ParseFloat(projectionData.CartDataOneYear[0].NabDaily, 64)
		projectionData.CagrOneYear = math.Round(((nabfloat-projectionData.CagrOneYear)/projectionData.CagrOneYear*100)*100) / 100

		return projectionData
	}
}
func getCartQuery(db *sql.DB, dir string, filename string, id string, time string, timeprojection string) []response.ProjectionCartData {
	var cartData []response.ProjectionCartData
	dat, err := ioutil.ReadFile(dir + filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	query := string(dat)
	query = fmt.Sprintf(query, id, timeprojection, time)
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
		cartData = append(cartData, response.ProjectionCartData{
			DateDaily:       temp1,
			DateDailyString: temp2,
			NabDaily:        temp3,
		})
	}
	return cartData
}
