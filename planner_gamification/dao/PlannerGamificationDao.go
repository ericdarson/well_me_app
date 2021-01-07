package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"planner_gamification_api/dbconnection"
	"planner_gamification_api/entity/response"

	_ "github.com/godror/godror"
)

type PlannerGamificationDao interface {
	GetDetailPlan(string) response.PlannerGamificationOutputSchema
}

type plannerGamificationDao struct {
}

func NewPlannerGamification() PlannerGamificationDao {
	return &plannerGamificationDao{}
}
func returnErrSimluation() response.PlannerGamificationOutputSchema {
	return response.PlannerGamificationOutputSchema{
		Nama:    "-1",
		Target:  -1,
		Amount:  -1,
		Gambar:  "-1",
		Puzzle:  "-1",
		DueDate: "-1",
	}
}

func (dao *plannerGamificationDao) GetDetailPlan(idjenis string) response.PlannerGamificationOutputSchema {
	var outputSchema response.PlannerGamificationOutputSchema
	conn := dbconnection.New()
	db := conn.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	dat, err := ioutil.ReadFile(dir + "/query/GetDetailPlan.query")
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	query := string(dat)
	query = fmt.Sprintf(query, idjenis)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return returnErrSimluation()
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&outputSchema.Nama, &outputSchema.Target, &outputSchema.Amount, &outputSchema.Gambar, &outputSchema.DueDate)
		outputSchema.Puzzle = outputSchema.Gambar[1:10]
		outputSchema.Gambar = outputSchema.Gambar[:1]
	}

	return outputSchema
}
