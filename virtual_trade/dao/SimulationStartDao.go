package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"virtual_trade_api/dbconnection"
	"virtual_trade_api/entity/response"

	_ "github.com/godror/godror"
)

type SimulationStartDao interface {
	StartSimulation(string, string) response.SimulationStartOutputSchema
}

type simulationStartDao struct {
}

func NewSimulationStart() SimulationStartDao {
	return &simulationStartDao{}
}
func returnErrSimulationStart() response.SimulationStartOutputSchema {
	return response.SimulationStartOutputSchema{
		JumlahUnit:      -1,
		StartingNab:     "-1",
		StartDate:       "-1",
		StartDateString: "-1",
	}
}

func (dao *simulationStartDao) StartSimulation(idproduk string, jumlahinvest string) response.SimulationStartOutputSchema {
	jumlahinvestfloat, _ := strconv.ParseFloat(jumlahinvest, 64)
	var simulationStartData response.SimulationStartOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrSimulationStart()
	}
	//validasi minumum pembelian
	dat, err := ioutil.ReadFile(dir + "/query/GetMinimumTransaction.query")
	if err != nil {
		fmt.Println(err)
		return returnErrSimulationStart()
	}
	query := string(dat)
	query = fmt.Sprintf(query, idproduk)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrSimulationStart()
	}
	defer rows.Close()
	boolErr := true
	for rows.Next() {
		boolErr = false
		var minimumTransaction float64
		rows.Scan(&minimumTransaction)
		if minimumTransaction > jumlahinvestfloat {
			return response.SimulationStartOutputSchema{
				JumlahUnit:      -2,
				StartingNab:     "-2",
				StartDate:       "-2",
				StartDateString: "-2",
			}
		}
	}
	if boolErr {
		returnErrSimulationStart()
	}

	dat, err = ioutil.ReadFile(dir + "/query/StartSimulation.query")
	if err != nil {
		fmt.Println(err)
		return returnErrSimulationStart()
	}
	query = string(dat)
	query = fmt.Sprintf(query, idproduk, idproduk)
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrSimulationStart()
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&simulationStartData.StartingNab, &simulationStartData.StartDate, &simulationStartData.StartDateString)
		startingnab, _ := strconv.ParseFloat(simulationStartData.StartingNab, 64)
		simulationStartData.JumlahUnit = jumlahinvestfloat / startingnab
	}

	return simulationStartData
}
