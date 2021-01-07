package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"virtual_trade_api/dbconnection"
	"virtual_trade_api/entity/response"

	_ "github.com/godror/godror"
)

type SimulationDao interface {
	GetSimulationResult(string, string) response.SimulationOutputSchema
}

type simulationDao struct {
}

func NewSimulation() SimulationDao {
	return &simulationDao{}
}
func returnErrSimluation() response.SimulationOutputSchema {
	return response.SimulationOutputSchema{
		DateSimulation:       "-1",
		DateSimulationString: "-1",
		NabSimulation:        "-1",
	}
}

func (dao *simulationDao) GetSimulationResult(idproduk string, simulationdate string) response.SimulationOutputSchema {
	var simulationData response.SimulationOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	dat, err := ioutil.ReadFile(dir + "/query/Simulation.query")
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	query := string(dat)
	query = fmt.Sprintf(query, idproduk, idproduk, simulationdate)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&simulationData.NabSimulation, &simulationData.DateSimulation, &simulationData.DateSimulationString)
	}

	return simulationData
}
