package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"planner_resiko_api/dbconnection"
	"planner_resiko_api/entity/response"

	_ "github.com/godror/godror"
)

type PlannerResikoDao interface {
	GetResiko(string) []response.PlannerResikoOutputSchema
}

type plannerResikoDao struct {
}

func NewPlannerResiko() PlannerResikoDao {
	return &plannerResikoDao{}
}
func returnErrResiko() []response.PlannerResikoOutputSchema {
	return []response.PlannerResikoOutputSchema{{
		IDJenis:    "-1",
		Nama:       "-1",
		Percentage: "-1",
	}}
}

func (dao *plannerResikoDao) GetResiko(bcaid string) []response.PlannerResikoOutputSchema {
	var outputSchema []response.PlannerResikoOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrResiko()
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetResikoByID.query")
	if err != nil {
		fmt.Println(err)
		return returnErrResiko()
	}
	query := string(dat)
	query = fmt.Sprintf(query, bcaid)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrResiko()
	}
	defer rows.Close()

	for rows.Next() {
		var single response.PlannerResikoOutputSchema
		rows.Scan(&single.IDJenis, &single.Nama, &single.Percentage)
		outputSchema = append(outputSchema, single)
	}

	return outputSchema
}
